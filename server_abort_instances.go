package later

import (
	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// AbortInstances aborts a bunch of instances
func (server *Server) AbortInstances(context context.Context, in *pb.AbortInstancesInput) (*pb.AbortInstancesOutput, error) {

	// Abort instances both locally and on the database
	instancesIDs, err := server.Machine.AbortInstances(in.GetTaskName(), in.GetParameters())
	if err != nil {
		return nil, err
	}

	out := &pb.AbortInstancesOutput{
		InstancesIds: instancesIDs,
		Error:        nil,
	}

	return out, nil

}
