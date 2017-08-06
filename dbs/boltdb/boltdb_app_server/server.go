package boltdb_app_server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/hippoai/later/dbs/boltdb/boltdb_app_server/_proto"
	"golang.org/x/net/context"
)

// Server wraps the database to serve its methods over gRPC and HTTP
type Server struct {
	Database *Database
}

// AbortInstance gRPC connector
func (server *Server) AbortInstance(ctx context.Context, in *pb.AbortInstanceInput) (*pb.AbortInstanceOutput, error) {

	err := server.Database.AbortInstance(in.GetTaskName(), in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	return &pb.AbortInstanceOutput{}, nil

}

// CreateInstance gRPC connector
func (server *Server) CreateInstance(ctx context.Context, in *pb.CreateInstanceInput) (*pb.CreateInstanceOutput, error) {

	executionTime, err := time.Parse(time.RFC3339, in.GetExecutionTime())
	if err != nil {
		return nil, err
	}

	instanceID, err := server.Database.CreateInstance(in.GetTaskName(), executionTime, in.GetParameters())
	if err != nil {
		return nil, err
	}

	return &pb.CreateInstanceOutput{
		InstanceId: instanceID,
	}, nil

}

// GetInstances gRPC connector
func (server *Server) GetInstances(ctx context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {

	start, err := time.Parse(time.RFC3339, in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, in.GetEnd())
	if err != nil {
		return nil, err
	}

	instances, err := server.Database.GetInstances(start, end)
	if err != nil {
		return nil, err
	}

	proto_instances := []*pb.Instance{}
	for _, instance := range instances {
		proto_instances = append(proto_instances, &pb.Instance{
			Id:            instance.ID,
			TaskName:      instance.TaskName,
			ExecutionTime: instance.ExecutionTime.Format(time.RFC3339),
			Parameters:    instance.Parameters,
		})
	}

	return &pb.GetInstancesOutput{
		Instances: proto_instances,
	}, nil

}

// GetLastPullTime gRPC connector
func (server *Server) GetLastPullTime(ctx context.Context, in *pb.GetLastPullTimeInput) (*pb.GetLastPullTimeOutput, error) {

	lastPullTime, err := server.Database.GetLastPullTime()
	if err != nil {
		return nil, err
	}

	// If it was not set yet
	if lastPullTime == nil {
		return &pb.GetLastPullTimeOutput{
			Time: nil,
		}, nil
	}

	return &pb.GetLastPullTimeOutput{
		Time: &pb.WrappedTime{
			Time: lastPullTime.Format(time.RFC3339),
		},
	}, nil

}

// MarkAsSuccessful gRPC connector
func (server *Server) MarkAsSuccessful(ctx context.Context, in *pb.MarkAsSuccessfulInput) (*pb.MarkAsSuccessfulOutput, error) {

	err := server.Database.MarkAsSuccessful(in.GetTaskName(), in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	return &pb.MarkAsSuccessfulOutput{}, nil

}

// MarkAsFailed gRPC connector
func (server *Server) MarkAsFailed(ctx context.Context, in *pb.MarkAsFailedInput) (*pb.MarkAsFailedOutput, error) {

	err := server.Database.MarkAsFailed(in.GetTaskName(), in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	return &pb.MarkAsFailedOutput{}, nil

}

// SetPullTime gRPC connector
func (server *Server) SetPullTime(ctx context.Context, in *pb.SetPullTimeInput) (*pb.SetPullTimeOutput, error) {

	t, err := time.Parse(time.RFC3339, in.GetTime())
	if err != nil {
		return nil, err
	}

	err = server.Database.SetPullTime(t)
	if err != nil {
		return nil, err
	}

	return &pb.SetPullTimeOutput{}, nil

}

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
