package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/cipher"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/config"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/data"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/jwt"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/oauth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const EXPIRES = 24 * time.Hour

type UserInfo struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

const COOKIE_STATE_NAME = "DONOTDEVELOPMYAPPRANDOMSTATE"
const COOKIE_JWT_NAME = "DONOTDEVELOPMYAPPJWT"

type myAuthServiceServer struct {
	userRepository data.UserRepository
	config         config.ConfigProvider
	cipherUtil     cipher.Cipher
	jwt            jwt.JwtProvider
	googleAuth     oauth.GoogleActions
}

func (s *myAuthServiceServer) init() {
	s.userRepository = data.NewUserRepository()
	s.cipherUtil = cipher.CipherUtil{}
	s.jwt = jwt.JwtImpl{}
	s.googleAuth = oauth.NewGoogleActions()
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	authServer := &myAuthServiceServer{config: config.OsEnv{}}
	authServer.init()
	defer authServer.userRepository.CloseConn()

	http.HandleFunc("/login", authServer.handleGoogleLogin)
	http.HandleFunc("/callback", authServer.oauthGoogleCallback)

	grpcServer := server.NewGRPC()

	auth.RegisterAuthServiceServer(grpcServer, authServer)

	go func() {
		log.Fatal(server.ServeGRPC(grpcServer))
		wg.Done()
	}()

	go func() {
		log.Fatal(server.ServeHTTPS())
		wg.Done()
	}()

	wg.Wait()
}

func (s *myAuthServiceServer) GetUser(ctx context.Context, token *auth.Token) (*auth.User, error) {
	userId, err := s.jwt.DecodeToken(token.Value, []byte(s.config.Get("ACCESS_SECRET")))
	return &auth.User{Id: userId}, err
}

func (s *myAuthServiceServer) generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(EXPIRES)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error creating random bytes: %s\n", err)
	}
	state := base64.URLEncoding.EncodeToString(b)

	encrypted, err := s.cipherUtil.Encrypt([]byte(s.config.Get("STATE_SECRET")), b)
	if err != nil {
		log.Fatalf("Error encrypting state: %s\n", err)
	}

	cookie := http.Cookie{Name: COOKIE_STATE_NAME, Value: base64.URLEncoding.EncodeToString(encrypted), Expires: expiration, HttpOnly: true, SameSite: http.SameSiteLaxMode}
	http.SetCookie(w, &cookie)

	return state
}

func (s *myAuthServiceServer) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := s.generateStateOauthCookie(w)
	url := s.googleAuth.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (s *myAuthServiceServer) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, err := r.Cookie(COOKIE_STATE_NAME)

	if err != nil {
		log.Printf("error obtaining state cookie: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	oauthStateEncrypted, err := base64.URLEncoding.DecodeString(oauthState.Value)
	if err != nil {
		log.Printf("cant decode state cookie, %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	decryptedState, err := s.cipherUtil.Decrypt([]byte(s.config.Get("STATE_SECRET")), oauthStateEncrypted)
	if err != nil {
		log.Printf("cant decrypt state cookie, %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if r.FormValue("state") != base64.URLEncoding.EncodeToString(decryptedState) {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	userData, err := s.googleAuth.GetUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var user UserInfo
	err = json.Unmarshal(userData, &user)
	if err != nil {
		fmt.Fprintf(w, "error decoding json: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if user.ID == "" {
		log.Println("bad response of google API, no user id in response")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var userDB data.User
	if err := s.userRepository.GetOrCreate(&userDB, &data.User{ID: user.ID, Name: "TOBEGENERATED"}); err != nil {
		fmt.Fprintf(w, "error saving user id: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := s.jwt.CreateToken(user.ID, []byte(s.config.Get("ACCESS_SECRET")), EXPIRES)
	if err != nil {
		fmt.Fprintf(w, "error while generating jwt: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	cookie := http.Cookie{Name: COOKIE_JWT_NAME, Value: token, Expires: time.Now().Add(EXPIRES), HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/"}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, s.config.Get("FRONT_URL"), http.StatusFound)

}
