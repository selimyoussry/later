package later

import (
	pb "github.com/hippoai/later/_proto"
	"golang.org/x/net/context"
)

// Stats returns the stats for this server
func (server *Server) Stats(context context.Context, in *pb.StatsInput) (*pb.StatsOutput, error) {

	// Calculate the number
	out := &pb.StatsOutput{
		NInMemory: server.Machine.GetNumberOfLocalInstances(),
	}

	return out, nil

}
