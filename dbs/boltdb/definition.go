package boltdb

import "github.com/hippoai/env"

type Database struct {
	gRPC_Address string
	Client       interface{}
}

// NewDatabase instanciates
func NewDatabase(gRPC_Address string) (*Database, error) {
	return &Database{
		gRPC_Address: gRPC_Address,
		Client:       nil,
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
