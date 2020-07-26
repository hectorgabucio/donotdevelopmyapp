package server

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

const HTTP_PORT = 8080
const GRPC_PORT = 8081

func ServeGRPC(server *grpc.Server) error {
	lis, err := net.Listen("tcp", port(GRPC_PORT))
	if err != nil {
		return err
	}
	return server.Serve(lis)
}

func ServeHTTP(mux *http.ServeMux) error {
	return http.ListenAndServe(port(HTTP_PORT), mux)
}

func port(port int) string {
	return fmt.Sprintf(":%d", port)
}
