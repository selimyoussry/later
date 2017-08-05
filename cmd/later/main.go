package main

import (
	"log"

	"github.com/hippoai/goutil"
	"github.com/hippoai/later"
	"github.com/hippoai/later/dbs/boltdb_single"
	"github.com/hippoai/later/tasks/bash"
	"github.com/hippoai/later/tasks/echo"
)

func main() {

	// List all tasks names
	tasks := []string{
		"echo", "bash",
	}

	// Use boltdb_single database, which runs on the same server as this node
	db, err := boltdb_single.NewDatabase(tasks)
	if err != nil {
		log.Fatal(err)
	}

	// Create a machine with default parameters
	machine := later.NewMachine(db, nil)

	// Register tasks
	err = machine.RegisterTasks(
		&bash.Task{},
		&echo.Task{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Listen on gRPC
	gRPC_server := later.NewServer(machine, "secret")
	go func() {
		for {
			err := gRPC_server.Run_gRPC()
			log.Printf("Error with gRPC server %s \n", goutil.Pretty(err))
		}
	}()

	// Run the machine
	err = machine.Loop()
	if err != nil {
		log.Fatal(err)
	}

}
