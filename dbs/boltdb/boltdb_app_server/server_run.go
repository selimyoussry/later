package boltdb_app_server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
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

// Run_HTTP_gRPC server
func (server *Server) Run_HTTP_gRPC() error {

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

	portStr := fmt.Sprintf(":%d", HTTP_gRPC_Server_Port)
	return http.ListenAndServe(portStr, mux)
}

// Run_HTTP serves the /export function
func (server *Server) Run_HTTP() error {

	backupHandleFunc := func(w http.ResponseWriter, req *http.Request) {
		err := server.Database.DB.View(func(tx *bolt.Tx) error {
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", `attachment; filename="later-export.db"`)
			w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
			_, err := tx.WriteTo(w)
			return err
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.HandleFunc("/export", backupHandleFunc)

	portStr := fmt.Sprintf(":%d", HTTP_Server_Port)
	return http.ListenAndServe(portStr, nil)

}
