package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
)

type randomHandler struct {
	client random.RandomServiceClient
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received new request: %s", r.URL.RequestURI())
		next.ServeHTTP(w, r)
		log.Printf("Sending response...")
	})
}

func (rh *randomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message, err := rh.client.GetRandom(context.Background(), &random.RandomInput{Max: 1000})
	if err != nil {
		log.Fatalf("error while saying hello to random micro %s", err)
	}

	fmt.Fprint(w, message.String())
}

func main() {
	conn, err := grpc.Dial(os.Getenv("RANDOM_MICRO_SERVICE_HOST")+":"+os.Getenv("RANDOM_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}
	defer conn.Close()
	client := random.NewRandomServiceClient(conn)

	rh := &randomHandler{client: client}
	randomHandler := http.HandlerFunc(rh.ServeHTTP)

	mux := http.NewServeMux()
	mux.Handle("/random", middlewareOne(randomHandler))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
