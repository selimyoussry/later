package boltdb_app_server

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/hippoai/goutil"
)

type Database struct {
	DB    *bolt.DB
	Tasks []string
}

// Open the database connection
func NewDatabase(tasks []string) (*Database, error) {

	// Open the database connection (it is a single file for BoltDB)
	db, err := bolt.Open(GetPath(), 0600, &bolt.Options{
		Timeout: 5 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	// Initialize the buckets
	err = Initialize(db, tasks)
	if err != nil {
		return nil, err
	}

	goutil.Log("[BoltDB] Started a database, stored in %s",
		GetPath(),
	)

	return &Database{
		DB:    db,
		Tasks: tasks,
	}, nil

}

// Close
func (database *Database) Close() error {
	return database.DB.Close()
}
