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
func (database *Database) AbortInstance(instanceID string) error {

	_, err := database.Client.AbortInstance(context.Background(), &pb.AbortInstanceInput{
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

// GetAborted gRPC translation
func (database *Database) GetAborted(start, end time.Time) ([]*structures.Instance, error) {

	instances := []*structures.Instance{}
	proto_instances, err := database.Client.GetAborted(context.Background(), &pb.GetInstancesInput{
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

// GetFailed gRPC translation
func (database *Database) GetFailed(start, end time.Time) ([]*structures.Instance, error) {

	instances := []*structures.Instance{}
	proto_instances, err := database.Client.GetFailed(context.Background(), &pb.GetInstancesInput{
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

// GetSuccessful gRPC translation
func (database *Database) GetSuccessful(start, end time.Time) ([]*structures.Instance, error) {

	instances := []*structures.Instance{}
	proto_instances, err := database.Client.GetSuccessful(context.Background(), &pb.GetInstancesInput{
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

// MarkAsSuccessful gRPC translation
func (database *Database) MarkAsSuccessful(instanceID string) error {

	_, err := database.Client.MarkAsSuccessful(context.Background(), &pb.MarkAsSuccessfulInput{
		InstanceId: instanceID,
	})
	if err != nil {
		return err
	}

	return nil
}

// MarkAsFailed gRPC translation
func (database *Database) MarkAsFailed(instanceID string) error {

	_, err := database.Client.MarkAsFailed(context.Background(), &pb.MarkAsFailedInput{
		InstanceId: instanceID,
	})
	if err != nil {
		return err
	}

	return nil
}
