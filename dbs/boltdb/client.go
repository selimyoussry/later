package boltdb

import (
	"time"

	"golang.org/x/net/context"

	pb "github.com/hippoai/later/dbs/boltdb/boltdb_app_server/_proto"
	"github.com/hippoai/later/structures"
	"google.golang.org/grpc"
)

// New_gRPC_Channel creates a connection and the client
func New_gRPC_Channel(gRPC_Address string) (*grpc.ClientConn, pb.LaterBoltDBClient, error) {

	// Open the grpc connection
	conn, err := grpc.Dial(gRPC_Address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewLaterBoltDBClient(conn)

	return conn, client, nil
}

// AbortInstance gRPC translation
func (database *Database) AbortInstance(taskName string, instanceID string) error {

	_, err := database.Client.AbortInstance(context.Background(), &pb.AbortInstanceInput{
		TaskName:   taskName,
		InstanceId: instanceID,
	})

	if err != nil {
		return err
	}

	return nil
}

// AbortInstance gRPC translation
func (database *Database) CreateInstance(taskName string, executionTime time.Time, parameters []byte) (string, error) {

	proto_instanceID, err := database.Client.CreateInstance(context.Background(), &pb.CreateInstanceInput{
		TaskName:      taskName,
		ExecutionTime: executionTime.Format(time.RFC3339),
		Parameters:    parameters,
	})

	if err != nil {
		return "", err
	}

	return proto_instanceID.GetInstanceId(), nil
}

// GetInstances gRPC translation
func (database *Database) GetInstances(start, end time.Time) ([]*structures.Instance, error) {

	instances := []*structures.Instance{}
	proto_instances, err := database.Client.GetInstances(context.Background(), &pb.GetInstancesInput{
		Start: start.Format(time.RFC3339),
		End:   end.Format(time.RFC3339),
	})
	if err != nil {
		return []*structures.Instance{}, err
	}

	for _, proto_instance := range proto_instances.GetInstances() {
		executionTime, err := time.Parse(time.RFC3339, proto_instance.GetExecutionTime())
		if err != nil {
			return []*structures.Instance{}, err
		}

		instance := &structures.Instance{
			ExecutionTime: executionTime,
			ID:            proto_instance.GetId(),
			TaskName:      proto_instance.GetTaskName(),
			Parameters:    proto_instance.GetParameters(),
		}
		instances = append(instances, instance)
	}

	return instances, nil
}

// GetLastPullTime gRPC translation
func (database *Database) GetLastPullTime() (*time.Time, error) {

	proto_t, err := database.Client.GetLastPullTime(context.Background(), &pb.GetLastPullTimeInput{})
	if err != nil {
		return nil, err
	}

	// Time not defined yet
	if proto_t.Time == nil {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, proto_t.GetTime().GetTime())
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// MarkAsSuccessful gRPC translation
func (database *Database) MarkAsSuccessful(taskName, instanceID string) error {

	_, err := database.Client.MarkAsSuccessful(context.Background(), &pb.MarkAsSuccessfulInput{
		TaskName:   taskName,
		InstanceId: instanceID,
	})
	if err != nil {
		return err
	}

	return nil
}

// MarkAsFailed gRPC translation
func (database *Database) MarkAsFailed(taskName, instanceID string) error {

	_, err := database.Client.MarkAsFailed(context.Background(), &pb.MarkAsFailedInput{
		TaskName:   taskName,
		InstanceId: instanceID,
	})
	if err != nil {
		return err
	}

	return nil
}

// SetPullTime gRPC translation
func (database *Database) SetPullTime(t time.Time) error {

	_, err := database.Client.SetPullTime(context.Background(), &pb.SetPullTimeInput{
		Time: t.Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	return nil
}
