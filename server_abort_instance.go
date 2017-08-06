package later

import (
	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// AbortInstances aborts a bunch of instances
func (server *Server) AbortInstance(context context.Context, in *pb.AbortInstanceInput) (*pb.AbortInstanceOutput, error) {

	// Abort instances both locally and on the database
	err := server.Machine.AbortInstance(in.GetInstanceId())
	if err != nil {
		return nil, err
	}

	out := &pb.AbortInstanceOutput{}

	return out, nil

}
