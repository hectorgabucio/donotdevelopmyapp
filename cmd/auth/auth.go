package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
)

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const COOKIE_STATE_NAME = "DONOTDEVELOPMYAPPRANDOMSTATE"
const COOKIE_JWT_NAME = "DONOTDEVELOPMYAPPJWT"
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

const EXPIRES = 24 * time.Hour

type myAuthServiceServer struct {
}

func (s *myAuthServiceServer) GetUser(ctx context.Context, token *auth.Token) (*auth.User, error) {
	userId, err := DecodeToken(token.Value)
	return &auth.User{Id: userId}, err
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleGoogleLogin)
	mux.HandleFunc("/callback", oauthGoogleCallback)

	grpcServer := grpc.NewServer()

	authServer := &myAuthServiceServer{}

	auth.RegisterAuthServiceServer(grpcServer, authServer)

	go func() {
		log.Fatal(server.ServeGRPC(grpcServer))
		wg.Done()
	}()

	go func() {
		log.Fatal(server.ServeHTTP(mux))
		wg.Done()
	}()

	wg.Wait()
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(EXPIRES)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error creating random bytes: %s\n", err)
	}
	state := base64.URLEncoding.EncodeToString(b)

	encrypted, err := Encrypt([]byte(os.Getenv("STATE_SECRET")), b)
	if err != nil {
		log.Fatalf("Error encrypting state: %s\n", err)
	}

	cookie := http.Cookie{Name: COOKIE_STATE_NAME, Value: base64.URLEncoding.EncodeToString(encrypted), Expires: expiration, HttpOnly: true, SameSite: http.SameSiteLaxMode}
	http.SetCookie(w, &cookie)

	return state
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
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

	decryptedState, err := Decrypt([]byte(os.Getenv("STATE_SECRET")), oauthStateEncrypted)
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

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var user UserInfo
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Fprintf(w, "error decoding json: %s\n", err)
		return
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		fmt.Fprintf(w, "error while generating jwt: %s\n", err)
		return
	}

	cookie := http.Cookie{Name: COOKIE_JWT_NAME, Value: token, Expires: time.Now().Add(EXPIRES), HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/"}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, os.Getenv("FRONT_URL"), 302)

}

func CreateToken(userid string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(EXPIRES).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	jwt, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	if !jwt.Valid {
		return "", fmt.Errorf("Token not valid")
	}

	for key, val := range claims {
		if key == "user_id" {
			return val.(string), nil
		}

	}

	return "", fmt.Errorf("Could not find user id claim in decoded token")

}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func Encrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
