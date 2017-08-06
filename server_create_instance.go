package later

import (
	"time"

	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// CreateInstance adds an instance to the database
func (server *Server) CreateInstance(context context.Context, in *pb.CreateInstanceInput) (*pb.CreateInstanceOutput, error) {

	// Parse the execution time
	executionTime, err := time.Parse(time.RFC3339, in.GetExecutionTime())
	if err != nil {
		return nil, err
	}

	// Create the instance
	instance, err := server.Machine.CreateInstance(
		in.GetTaskName(),
		executionTime,
		in.GetParameters(),
	)
	if err != nil {
		return nil, err
	}

	out := &pb.CreateInstanceOutput{
		InstanceId: instance.ID,
	}

	return out, nil

}
