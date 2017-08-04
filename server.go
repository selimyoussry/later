package later

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hippoai/env"
	"google.golang.org/grpc"

	pb "github.com/hippoai/later/_proto"
)

type Server struct {
	Machine   *Machine
	SecretKey string
}

// NewServer instanciates
func NewServer(machine *Machine, secretKey string) *Server {
	return &Server{
		Machine:   machine,
		SecretKey: secretKey,
	}
}

// NewServerFromEnv instanciates from env variables
func NewServerFromEnv(machine *Machine) (*Server, error) {

	parsed, err := env.Parse(Env_SecretKey)
	if err != nil {
		return nil, err
	}

	return NewServer(machine, parsed[Env_SecretKey]), nil

}

// Run the gRPC and HTTP servers
func (server *Server) Run_gRPC() error {

	portStr := fmt.Sprintf(":%d", gRPC_Server_Port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMutationsServer(grpcServer, server)

	log.Printf("Running gRPC server on port %d \n", gRPC_Server_Port)
	return grpcServer.Serve(lis)
}

// Run_HTTP server
func (server *Server) Run_HTTP() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	gRPC_Endpoint := fmt.Sprintf("localhost:%d", gRPC_Server_Port)
	err := pb.RegisterMutationsHandlerFromEndpoint(ctx, mux, gRPC_Endpoint, opts)
	if err != nil {
		return err
	}

	portStr := fmt.Sprintf(":%d", HTTP_Server_Port)
	return http.ListenAndServe(portStr, mux)
}
