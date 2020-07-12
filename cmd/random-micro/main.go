package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/chat"
	"google.golang.org/grpc"
)

type chatHandler struct{}

func (c chatHandler) SayHello(ctx context.Context, msg *chat.Message) (*chat.Message, error) {
	log.Printf("Received message: %s", msg.String())
	return &chat.Message{Body: "HELLOOOOO"}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatHandler{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
