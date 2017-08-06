package boltdb

import (
	"github.com/hippoai/env"
	pb "github.com/hippoai/later/dbs/boltdb/boltdb_app_server/_proto"
	"google.golang.org/grpc"
)

// Database implements a "later" package Database interface
type Database struct {
	gRPC_Address string
	connection   *grpc.ClientConn
	Client       pb.LaterBoltDBClient
}

// NewDatabase instanciates
func NewDatabase(gRPC_Address string) (*Database, error) {

	// Create a new gRPC connection and client
	conn, client, err := New_gRPC_Channel(gRPC_Address)
	if err != nil {
		return nil, err
	}

	return &Database{
		gRPC_Address: gRPC_Address,
		connection:   conn,
		Client:       client,
	}, nil
}

// NewDatabaseFromEnv
func NewDatabaseFromEnv() (*Database, error) {

	parsed, err := env.Parse(Env_gRPC_Address)
	if err != nil {
		return nil, err
	}

	return NewDatabase(parsed[Env_gRPC_Address])

}

// Close closes the database connection
func (database *Database) Close() error {

	return database.connection.Close()

}
