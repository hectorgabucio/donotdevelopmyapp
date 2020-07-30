package grpctest

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener

const bufSize = 1024 * 1024

func InitServer() *grpc.Server {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return s
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Dialer() *grpc.ClientConn {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}
