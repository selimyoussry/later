package later

import (
	"time"

	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// getter returns a bunch of instances between a start and end time
func (server *Server) getter(context context.Context, in *pb.GetInstancesInput, status string) (*pb.GetInstancesOutput, error) {

	// Parse start and end time
	start, err := time.Parse(time.RFC3339, in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, in.GetEnd())
	if err != nil {
		return nil, err
	}

	instances, err := server.Machine.getter(start, end, status)
	if err != nil {
		return nil, err
	}

	// Convert them to the proto format
	proto_instances := []*pb.Instance{}
	for _, instance := range instances {
		proto_instances = append(proto_instances, &pb.Instance{
			Id:            instance.ID,
			TaskName:      instance.TaskName,
			ExecutionTime: instance.ExecutionTime.Format(time.RFC3339),
			Parameters:    instance.Parameters,
		})
	}

	// The protobuf output
	out := &pb.GetInstancesOutput{
		Instances: proto_instances,
	}

	return out, nil

}

// GetInstances returns a bunch of instances between a start and end time
func (server *Server) GetInstances(context context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.getter(context, in, STATUS_PENDING)
}

// GetFailed returns a bunch of instances between a start and end time
func (server *Server) GetFailed(context context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.getter(context, in, STATUS_FAILED)
}

// GetAborted returns a bunch of instances between a start and end time
func (server *Server) GetAborted(context context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.getter(context, in, STATUS_ABORTED)
}

// GetSuccessful returns a bunch of instances between a start and end time
func (server *Server) GetSuccessful(context context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {
	return server.getter(context, in, STATUS_SUCCESSFUL)
}
