package boltdb_app_server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/hippoai/later/dbs/boltdb/boltdb_app_server/_proto"
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Run the gRPC server
func (server *Server) Run_gRPC() error {

	portStr := fmt.Sprintf(":%d", gRPC_Server_Port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLaterBoltDBServer(grpcServer, server)

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
	err := pb.RegisterLaterBoltDBHandlerFromEndpoint(ctx, mux, gRPC_Endpoint, opts)
	if err != nil {
		return err
	}

	portStr := fmt.Sprintf(":%d", HTTP_Server_Port)
	return http.ListenAndServe(portStr, mux)
}
