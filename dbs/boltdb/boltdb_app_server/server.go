package boltdb_app_server

import (
	"time"

	pb "github.com/hippoai/later/dbs/boltdb/boltdb_app_server/_proto"
	"golang.org/x/net/context"
)

// Server wraps the database to serve its methods over gRPC and HTTP
type Server struct {
	Database *Database
}

// AbortInstance gRPC connector
func (server *Server) AbortInstance(ctx context.Context, in *pb.AbortInstanceInput) (*pb.AbortInstanceOutput, error) {

	err := server.Database.AbortInstance(in.GetInstanceId())
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
func (server *Server) get(ctx context.Context, in *pb.GetInstancesInput, bucketName string) (*pb.GetInstancesOutput, error) {

	start, err := time.Parse(time.RFC3339, in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, in.GetEnd())
	if err != nil {
		return nil, err
	}

	instances, err := server.Database.get(start, end, bucketName)
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

// GetInstances
func (server *Server) GetInstances(ctx context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.get(ctx, in, BUCKET_PENDING)
}

// GetAborted
func (server *Server) GetAborted(ctx context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.get(ctx, in, BUCKET_ABORTED)
}

// GetSuccessful
func (server *Server) GetSuccessful(ctx context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.get(ctx, in, BUCKET_SUCCESSFUL)
}

// GetFailed
func (server *Server) GetFailed(ctx context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.get(ctx, in, BUCKET_FAILED)
}

// MarkAsSuccessful gRPC connector
func (server *Server) MarkAsSuccessful(ctx context.Context, in *pb.MarkAsSuccessfulInput) (*pb.MarkAsSuccessfulOutput, error) {

	err := server.Database.MarkAsSuccessful(in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	return &pb.MarkAsSuccessfulOutput{}, nil

}

// MarkAsFailed gRPC connector
func (server *Server) MarkAsFailed(ctx context.Context, in *pb.MarkAsFailedInput) (*pb.MarkAsFailedOutput, error) {

	err := server.Database.MarkAsFailed(in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	return &pb.MarkAsFailedOutput{}, nil

}
