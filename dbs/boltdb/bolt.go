package boltdb

import (
	"time"

	"github.com/boltdb/bolt"
)

type Database struct {
	DB    *bolt.DB
	Tasks []string
}

// Open the database connection
func NewDatabase(tasks []string) (*Database, error) {

	db, err := bolt.Open("later.db", 0600, &bolt.Options{
		Timeout: 5 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	// Create the buckets for each task
	for _, taskName := range tasks {
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(bucket(taskName))
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return &Database{
		DB:    db,
		Tasks: tasks,
	}, nil

}

// Close
func (database *Database) Close() error {
	return database.DB.Close()
}
