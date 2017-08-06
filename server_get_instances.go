package later

import (
	"time"

	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// GetInstances returns a bunch of instances between a start and end time
func (server *Server) GetInstances(context context.Context, in *pb.GetInstancesInput) (*pb.GetInstancesOutput, error) {

	// Parse start and end time
	start, err := time.Parse(time.RFC3339, in.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, in.GetEnd())
	if err != nil {
		return nil, err
	}

	// Get the instances from the database during this timeframe
	instances, err := server.Machine.Database.GetInstances(start, end)
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
