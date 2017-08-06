package boltdb_app_server

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hippoai/goutil"
	"github.com/hippoai/later/structures"
	laterutil "github.com/hippoai/later/util"
)

type KV struct {
	K []byte
	V []byte
}

// GetInstances returns all the instances scheduled between start and end time
func (database *Database) get(start, end time.Time, bucketName string) ([]*structures.Instance, error) {

	this_bucket_kv := []*KV{}

	// Execute the transaction and store the instance ID outside the scope if successful
	err := database.DB.View(func(tx *bolt.Tx) error {

		min := []byte(laterutil.TimeToString(start))
		max := []byte(laterutil.TimeToString(end))

		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = cursor.Next() {
			this_bucket_kv = append(this_bucket_kv, &KV{
				K: k,
				V: v,
			})
		}

		return nil
	})
	if err != nil {
		return []*structures.Instance{}, err
	}

	// Now unmarshal them into structures.Instance
	instances := []*structures.Instance{}

	errorsForKeys := []string{}
	for _, kv := range this_bucket_kv {

		// Get the key in Boltdb
		key := string(kv.K)

		// Get the instance stored in the database
		// And if we can't unmarshal it, return the error
		var instance structures.Instance
		err := json.Unmarshal(kv.V, &instance)
		if err != nil {
			errorsForKeys = append(errorsForKeys, key)
			break
		}

		instances = append(instances, &instance)

	}

	// Print the errors
	if len(errorsForKeys) > 0 {
		log.Printf("Errors unmarshaling in %s \n", goutil.Stringify(errorsForKeys))
	}

	// Still return the instances we could find, the others are corrupted.

	return instances, nil

}

// GetInstances returns the pending commands
func (database *Database) GetInstances(start, end time.Time) ([]*structures.Instance, error) {
	return database.get(start, end, BUCKET_PENDING)
}

// GetAborted
func (database *Database) GetAborted(start, end time.Time) ([]*structures.Instance, error) {
	return database.get(start, end, BUCKET_ABORTED)
}

// GetSuccessful
func (database *Database) GetSuccessful(start, end time.Time) ([]*structures.Instance, error) {
	return database.get(start, end, BUCKET_SUCCESSFUL)
}

// GetFailed
func (database *Database) GetFailed(start, end time.Time) ([]*structures.Instance, error) {
	return database.get(start, end, BUCKET_FAILED)
}
