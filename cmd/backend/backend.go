package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
)

const COOKIE_JWT_NAME = "DONOTDEVELOPMYAPPJWT"

type app struct {
	randomClient    random.RandomServiceClient
	characterClient character.CharacterServiceClient
	authClient      auth.AuthServiceClient
}

type characterJson struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("FRONT_URL"))
		w.Header().Add("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received new request: %s", r.URL.RequestURI())
		next.ServeHTTP(w, r)
		log.Printf("Sending response...")
	})
}

func (app *app) securedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("This request needs to be authenticated")
		authCookie, err := r.Cookie(COOKIE_JWT_NAME)
		if err != nil {
			log.Println("No cookie found on secure request, aborting")
			http.Error(w, "Not authorized", 401)
			return
		}
		user, err := app.authClient.GetUser(context.Background(), &auth.Token{Value: authCookie.Value})
		if err != nil {
			log.Println("Error validating auth cookie, aborting", err)
			http.Error(w, "Not authorized", 401)
			return
		}
		log.Println("Authorized, user is", user)
		next.ServeHTTP(w, r)
	})
}

func (rh *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message, err := rh.randomClient.GetRandom(context.Background(), &random.RandomInput{Max: 1000})
	if err != nil {
		log.Fatalf("error while saying hello to random micro %s", err)
	}

	numberStr := strconv.FormatUint(message.Number, 10)
	character, err := rh.characterClient.GetCharacter(context.Background(), &character.Input{Number: numberStr})
	if err != nil {
		log.Fatalf("Error while getting random character %s", err)
	}

	if character.Name == "" {
		w.WriteHeader(404)
		return
	}

	mappedResponse := &characterJson{Id: strconv.FormatInt(int64(character.Id), 10), Name: character.Name, Image: character.Image}

	encoded, err := json.Marshal(mappedResponse)
	if err != nil {
		log.Fatalf("Error encoding to json character %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(encoded))
}

func main() {

	connRandom, err := server.EstablishGRPCConn(os.Getenv("RANDOM_MICRO_SERVICE_HOST") + ":" + os.Getenv("RANDOM_MICRO_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connRandom.Close()
	randomClient := random.NewRandomServiceClient(connRandom)

	connCharacter, err := server.EstablishGRPCConn(os.Getenv("CHARACTER_MICRO_SERVICE_HOST") + ":" + os.Getenv("CHARACTER_MICRO_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connCharacter.Close()
	characterClient := character.NewCharacterServiceClient(connCharacter)

	connAuth, err := server.EstablishGRPCConn(os.Getenv("AUTH_MICRO_SERVICE_HOST") + ":" + os.Getenv("AUTH_MICRO_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connAuth.Close()
	authClient := auth.NewAuthServiceClient(connAuth)

	app := &app{randomClient: randomClient, characterClient: characterClient, authClient: authClient}
	handler := http.HandlerFunc(app.ServeHTTP)

	mux := http.NewServeMux()
	mux.Handle("/random", corsMiddleware(app.securedMiddleware((logMiddleware(handler)))))
	log.Fatal(server.ServeHTTP(mux))
}