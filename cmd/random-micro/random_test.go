package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	random.RegisterRandomServiceServer(s, &randomHandler{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetRandomNumber(t *testing.T) {

	tests := []struct {
		maxInput     int64
		errorMessage string
	}{
		{10, ""},
		{-1, "rpc error: code = Unknown desc = Max input cant be 0 or lesser than 0"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.maxInput, tt.errorMessage)
		t.Run(testname, func(t *testing.T) {

			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := random.NewRandomServiceClient(conn)
			_, err = client.GetRandom(ctx, &random.RandomInput{Max: tt.maxInput})
			if err != nil && tt.errorMessage != err.Error() {
				t.Fatalf("TestGetRandomNumber failed: %v", err)
			}

		})
	}

}
