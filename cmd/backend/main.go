package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type randomHandler struct {
	client random.RandomServiceClient
}

func (rh *randomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("New request: %s", r.URL.String())
	message, err := rh.client.GetRandom(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while saying hello to random micro %s", err)
	}

	log.Printf("Getting response from random micro: %s", message.String())

	fmt.Fprint(w, message.String())
}

func main() {
	mux := http.NewServeMux()

	conn, err := grpc.Dial(os.Getenv("RANDOM_MICRO_SERVICE_HOST")+":"+os.Getenv("RANDOM_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}

	defer conn.Close()
	client := random.NewRandomServiceClient(conn)

	rh := &randomHandler{client: client}

	mux.Handle("/random", rh)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
