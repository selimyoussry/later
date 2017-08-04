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

	// Store the instance in the database
	instanceID, err := server.Machine.Database.CreateInstance(in.GetTaskName(), executionTime, in.GetParameters())
	if err != nil {
		return nil, err
	}

	// Pull it locally if it's in the same timeframe
	timeframeEnd := time.Now().Add(server.Machine.Parameters.TimeAhead)
	if executionTime.Before(timeframeEnd) {

		instance := &Instance{
			ExecutionTime: executionTime,
			ID:            instanceID,
			Parameters:    in.GetParameters(),
			TaskName:      in.GetTaskName(),
		}

		server.Machine.StartInstance(instance)

	}

	out := &pb.CreateInstanceOutput{
		InstanceId: instanceID,
		Error:      nil,
	}

	return out, nil

}
