package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

const GRPC_PORT = 8081

func ServeGRPC(server *grpc.Server) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", GRPC_PORT))
	if err != nil {
		return err
	}
	return server.Serve(lis)
}
