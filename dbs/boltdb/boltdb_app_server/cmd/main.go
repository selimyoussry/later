package main

import (
	"log"

	"github.com/hippoai/later/dbs/boltdb/boltdb_app_server"
)

func main() {

	// Create the database
	db, err := boltdb_app_server.NewDatabase(
		[]string{"echo", "bash"},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Server over gRPC and HTTP
	server := &boltdb_app_server.Server{
		Database: db,
	}

	// Serve gRPC
	go func() {
		err := server.Run_gRPC()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Server HTTP
	go func() {
		err := server.Run_HTTP()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait forever
	select {}

}
