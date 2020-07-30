package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
)

type randomHandler struct{}

func (c randomHandler) GetRandom(ctx context.Context, input *random.RandomInput) (*random.RandomNumber, error) {
	if input.Max <= 0 {
		return nil, fmt.Errorf("Max input cant be 0 or lesser than 0")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(input.Max))
	if err != nil {
		return nil, err
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
