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

	db, err := bolt.Open(GetPath(), 0600, &bolt.Options{
		Timeout: 5 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	// Initialize the buckets
	err = db.Update(func(tx *bolt.Tx) error {

		var err error

		// Create the buckets for each task
		for _, taskName := range tasks {
			_, err = tx.CreateBucketIfNotExists(bucket(taskName))
			if err != nil {
				return err
			}
		}

		// Create the aborted bucket
		_, err = tx.CreateBucketIfNotExists([]byte(BUCKET_ABORTED))
		if err != nil {
			return err
		}

		// Create the successful bucket
		_, err = tx.CreateBucketIfNotExists([]byte(BUCKET_SUCCESSFUL))
		if err != nil {
			return err
		}

		// Create the failed bucket
		_, err = tx.CreateBucketIfNotExists([]byte(BUCKET_FAILED))
		if err != nil {
			return err
		}

		// Create the metadata bucket
		_, err = tx.CreateBucketIfNotExists([]byte(BUCKET_METADATA))
		if err != nil {
			return err
		}

		return nil

	})
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
