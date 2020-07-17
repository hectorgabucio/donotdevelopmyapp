package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/character"
	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
)

type app struct {
	randomClient    random.RandomServiceClient
	characterClient character.CharacterServiceClient
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received new request: %s", r.URL.RequestURI())
		next.ServeHTTP(w, r)
		log.Printf("Sending response...")
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

	app := &app{randomClient: randomClient, characterClient: characterClient}
	handler := http.HandlerFunc(app.ServeHTTP)

	mux := http.NewServeMux()
	mux.Handle("/random", middlewareOne(handler))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
