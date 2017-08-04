package boltdb

import (
	"bytes"
	"encoding/json"
	"sync"
	"time"

	"golang.org/x/sync/syncmap"

	"github.com/boltdb/bolt"
	"github.com/hippoai/later/laterutil"
)

type Instance struct{}

// GetInstances returns all the instances scheduled between start and end time
func (database *Database) GetInstances(start, end time.Time) ([]*Instance, error) {

	// Execute the transaction and store the instance ID outside the scope if successful
	err := database.DB.Update(func(tx *bolt.Tx) error {

		all := syncmap.Map{}

		min := laterutil.TimeToString(start)
		max := laterutil.TimeToString(end)

		// Iterate through all buckets
		var wg sync.WaitGroup
		for _, taskName := range database.Tasks {
			wg.Add(1)

			go func(taskName string) {

				defer wg.Done()
				cursor := tx.Bucket(bucket(taskName)).Cursor()
				instances := []*Instance{}

				for k, v := cursor.Seek(); k != nil && bytes.Compare(k, max) <= 0; k, v := cursor.Next() {

				}

			}(taskName)

		}

		wg.Wait()
		// Find the bucket
		b := tx.Bucket(bucket(name))

		// Make the key and value
		instanceID = MakeID(executionTime)
		value := &Value{
			ExecutionTime: executionTime.Format(time.RFC3339),
			ID:            instanceID,
			Parameters:    parameters,
		}
		valueBytes, err := json.Marshal(value)

		// Set the value
		err = b.Put([]byte(instanceID), valueBytes)
		if err != nil {
			return err
		}

		return nil
	})

	bytes.Compare(a, b)

	return instanceID, err

}
