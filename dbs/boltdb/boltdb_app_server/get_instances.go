package boltdb_app_server

import (
	"bytes"
	"encoding/json"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/syncmap"

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
func (database *Database) GetInstances(start, end time.Time) ([]*structures.Instance, error) {

	all := syncmap.Map{}

	// Execute the transaction and store the instance ID outside the scope if successful
	err := database.DB.View(func(tx *bolt.Tx) error {

		min := []byte(laterutil.TimeToString(start))
		max := []byte(laterutil.TimeToString(end))

		// Iterate through all buckets concurrently to gather the instances
		var wg sync.WaitGroup
		for _, taskName := range database.Tasks {
			wg.Add(1)

			go func(taskName string) {

				defer wg.Done()
				cursor := tx.Bucket(bucket(taskName)).Cursor()
				this_bucket_kv := []*KV{}

				for k, v := cursor.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = cursor.Next() {
					this_bucket_kv = append(this_bucket_kv, &KV{
						K: k,
						V: v,
					})
				}

				all.Store(taskName, this_bucket_kv)

			}(taskName)

		}

		wg.Wait()

		return nil
	})
	if err != nil {
		return []*structures.Instance{}, err
	}

	// Now unmarshal them into structures.Instance
	instances := []*structures.Instance{}

	errorsForKeys := []string{}
	all.Range(func(k interface{}, vItf interface{}) bool {
		kvs := vItf.([]*KV)

		for _, kv := range kvs {

			// Get the key in Boltdb
			key := string(kv.K)

			// Get the instance stored in the database
			// And if we can't unmarshal it, return the error
			var instance structures.Instance
			err := json.Unmarshal(kv.V, &instance)
			if err != nil {
				errorsForKeys = append(errorsForKeys, key)
				return true
			}

			instances = append(instances, &instance)

		}

		return true
	})

	// Print the errors
	if len(errorsForKeys) > 0 {
		log.Printf("Errors unmarshaling in %s \n", goutil.Stringify(errorsForKeys))
	}

	// Still return the instances we could find, the others are corrupted.

	return instances, nil

}
