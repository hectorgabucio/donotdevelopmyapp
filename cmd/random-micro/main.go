package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
)

type randomHandler struct{}

func (c randomHandler) GetRandom(ctx context.Context, empty *empty.Empty) (*random.RandomNumber, error) {
	return &random.RandomNumber{Number: 1}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := randomHandler{}

	grpcServer := grpc.NewServer()

	random.RegisterRandomServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
