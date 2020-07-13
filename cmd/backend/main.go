package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/chat"
	"google.golang.org/grpc"
)

type randomHandler struct {
	client chat.ChatServiceClient
}

func (rh *randomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message, err := rh.client.SayHello(context.Background(), &chat.Message{Body: "hola"})
	if err != nil {
		log.Fatalf("error while saying hello to random micro %s", err)
	}

	log.Printf("Getting response from random micro: %s", message.String())

	fmt.Fprintf(w, message.Body)
}

func main() {
	mux := http.NewServeMux()

	conn, err := grpc.Dial(os.Getenv("RANDOM_MICRO_SERVICE_HOST")+":"+os.Getenv("RANDOM_MICRO_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial grpc: %s", err)
	}

	defer conn.Close()
	client := chat.NewChatServiceClient(conn)

	rh := &randomHandler{client: client}

	mux.Handle("/", rh)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
