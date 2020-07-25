package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/pkg/character"
	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
)

const COOKIE_JWT_NAME = "DONOTDEVELOPMYAPPJWT"

type app struct {
	randomClient    random.RandomServiceClient
	characterClient character.CharacterServiceClient
	authClient      auth.AuthServiceClient
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

	fmt.Fprint(w, character.String())
}

func main() {

	connRandom, err := grpc.Dial(os.Getenv("RANDOM_MICRO_SERVICE_HOST")+":"+os.Getenv("RANDOM_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connRandom.Close()
	randomClient := random.NewRandomServiceClient(connRandom)

	connCharacter, err := grpc.Dial(os.Getenv("CHARACTER_MICRO_SERVICE_HOST")+":"+os.Getenv("CHARACTER_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connCharacter.Close()
	characterClient := character.NewCharacterServiceClient(connCharacter)

	connAuth, err := grpc.Dial(os.Getenv("AUTH_MICRO_SERVICE_HOST")+":"+os.Getenv("AUTH_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer connAuth.Close()
	authClient := auth.NewAuthServiceClient(connAuth)

	app := &app{randomClient: randomClient, characterClient: characterClient, authClient: authClient}
	handler := http.HandlerFunc(app.ServeHTTP)

	mux := http.NewServeMux()
	mux.Handle("/random", corsMiddleware(app.securedMiddleware((logMiddleware(handler)))))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
