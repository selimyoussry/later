package main

import (
	"log"

	"github.com/hippoai/goutil"
	"github.com/hippoai/later"
	"github.com/hippoai/later/dbs/boltdb"
	"github.com/hippoai/later/tasks/bash"
	"github.com/hippoai/later/tasks/echo"
)

func main() {

	// Use boltdb database, which runs on the same server as this node
	db, err := boltdb.NewDatabaseFromEnv()
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

	// Serve on HTTP
	go func() {
		for {
			err := gRPC_server.Run_HTTP()
			log.Printf("Error with HTTP server %s \n", err)
		}
	}()

	// Run the machine
	err = machine.Loop()
	if err != nil {
		log.Fatal(err)
	}

}
