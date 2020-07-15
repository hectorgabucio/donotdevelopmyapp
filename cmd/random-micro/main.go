package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/hectorgabucio/donotdevelopmyapp/pkg/random"
	"google.golang.org/grpc"
)

type randomHandler struct{}

func (c randomHandler) GetRandom(ctx context.Context, input *random.RandomInput) (*random.RandomNumber, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(input.Max))
	if err != nil {
		log.Fatalf("Error while calculating secure random number: %s", err)
	}
	return &random.RandomNumber{Number: n.Uint64()}, nil
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
