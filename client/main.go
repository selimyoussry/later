package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hippoai/goutil"
	pb "github.com/hippoai/later/_proto"
	"github.com/hippoai/later/tasks/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func get_instances() {

	// Create gRPC connection
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLaterClient(conn)

	out, err := client.GetInstances(context.Background(), &pb.GetInstancesInput{
		Start: time.Now().UTC().Add(-10 * time.Minute).Format(time.RFC3339),
		End:   time.Now().UTC().Add(10 * time.Minute).Format(time.RFC3339),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(goutil.Pretty(out.GetInstances()))

}

func add_instance() {

	// Create gRPC connection
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLaterClient(conn)

	msg := "Hello world!"
	parameters := &echo.Parameters{
		Message: &msg,
	}
	b, _ := json.Marshal(parameters)

	out, err := client.CreateInstance(context.Background(), &pb.CreateInstanceInput{
		TaskName:      "echo",
		ExecutionTime: time.Now().Add(20 * time.Second).Format(time.RFC3339),
		Parameters:    b,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(goutil.Pretty(out))

}

func abort_instance(instanceID string) {

	// Create gRPC connection
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLaterClient(conn)

	out, err := client.AbortInstance(context.Background(), &pb.AbortInstanceInput{
		TaskName:   "echo",
		InstanceId: instanceID,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(goutil.Pretty(out))

}

func stats() {

	// Create gRPC connection
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLaterClient(conn)

	out, err := client.Stats(context.Background(), &pb.StatsInput{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(goutil.Pretty(out.GetNInMemory()))

}
func main() {

	add_instance()
	get_instances()
	// abort_instance("2017-08-05T19:28:38-04:00.4596768c-9a1d-4640-a54e-58b1ce9778ce")
	stats()

}
