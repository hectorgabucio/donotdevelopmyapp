package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/test/grpctest"
)

func TestGetRandomNumber(t *testing.T) {

	tests := []struct {
		maxInput     int64
		errorMessage string
	}{
		{10, ""},
		{-1, "rpc error: code = Unknown desc = Max input cant be 0 or lesser than 0"},
	}

	s := grpctest.InitServer()
	defer s.GracefulStop()
	random.RegisterRandomServiceServer(s, &randomHandler{})

	conn := grpctest.Dialer()
	defer conn.Close()

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.maxInput, tt.errorMessage)
		t.Run(testname, func(t *testing.T) {
			ctx := context.Background()
			client := random.NewRandomServiceClient(conn)
			_, err := client.GetRandom(ctx, &random.RandomInput{Max: tt.maxInput})
			if err != nil && tt.errorMessage != err.Error() {
				t.Fatalf("TestGetRandomNumber failed: %v", err)
			}

		})
	}

}