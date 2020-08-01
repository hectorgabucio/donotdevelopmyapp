package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/test/grpctest"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomNumber(t *testing.T) {

	tests := []struct {
		maxInput     int64
		errorMessage string
	}{
		{10, ""},
		{-1, "rpc error: code = Unknown desc = Max input cant be 0 or lesser than 0"},
	}
	init := make(chan bool)
	s := grpctest.InitServer(init)
	defer s.GracefulStop()
	random.RegisterRandomServiceServer(s, &randomHandler{})
	init <- true

	conn := grpctest.Dialer()
	defer conn.Close()

	assert := assert.New(t)

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.maxInput, tt.errorMessage)
		t.Run(testname, func(t *testing.T) {
			ctx := context.Background()
			client := random.NewRandomServiceClient(conn)
			_, err := client.GetRandom(ctx, &random.RandomInput{Max: tt.maxInput})
			if err != nil {
				assert.Equal(tt.errorMessage, err.Error(), "It should be equal")
			}

		})
	}

}
