package boltdb_app_server

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hippoai/later/structures"
)

// CreateInstance creates a new instance in the database
func (database *Database) CreateInstance(taskName string, executionTime time.Time, parameters []byte) (string, error) {

	var instanceID string

	// Execute the transaction and store the instance ID outside the scope if successful
	err := database.DB.Update(func(tx *bolt.Tx) error {

		// Find the bucket
		b := tx.Bucket([]byte(BUCKET_PENDING))

		// Make sure it's UTC
		executionTime = executionTime.UTC()

		// Make the key and value
		instanceID = MakeID(executionTime, taskName)
		value := &structures.Instance{
			ExecutionTime: executionTime,
			ID:            instanceID,
			Parameters:    parameters,
			TaskName:      taskName,
		}
		valueBytes, err := json.Marshal(value)

		// Set the value
		err = b.Put([]byte(instanceID), valueBytes)
		if err != nil {
			return err
		}

		return nil
	})

	return instanceID, err

}
