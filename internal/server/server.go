package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const HTTPS_PORT = 8080
const GRPC_PORT = 8081

const TLS_PEM_PATH = "./tls/tls.crt"
const TLS_KEY_PATH = "./tls/tls.key"

func ServeGRPC(server *grpc.Server) error {
	lis, err := net.Listen("tcp", port(GRPC_PORT))
	if err != nil {
		return err
	}
	return server.Serve(lis)
}

func ServeHTTPS() error {
	return http.ListenAndServeTLS(port(HTTPS_PORT), TLS_PEM_PATH, TLS_KEY_PATH, nil)
}

func EstablishGRPCConn(addr string) (*grpc.ClientConn, error) {
	log.Println("Connecting to grpc addr:", addr)
	return grpc.Dial(addr, grpcClientCredentials())

}

func NewGRPC() *grpc.Server {
	return grpc.NewServer(grpcServerCredentials())
}

func grpcClientCredentials() grpc.DialOption {
	creds, err := credentials.NewClientTLSFromFile(TLS_PEM_PATH, "")
	if err != nil {
		log.Fatalf("could not process the credentials: %v", err)
	}
	return grpc.WithTransportCredentials(creds)
}

func grpcServerCredentials() grpc.ServerOption {
	creds, err := credentials.NewServerTLSFromFile(TLS_PEM_PATH, TLS_KEY_PATH)
	if err != nil {
		log.Fatalf("Failed to setup TLS: %v", err)
	}
	return grpc.Creds(creds)
}

func port(port int) string {
	return fmt.Sprintf(":%d", port)
}
