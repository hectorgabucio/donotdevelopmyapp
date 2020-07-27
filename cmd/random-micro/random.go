package main

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
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
	s := randomHandler{}
	grpcServer := server.NewGRPC()
	random.RegisterRandomServiceServer(grpcServer, &s)
	if err := server.ServeGRPC(grpcServer); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
