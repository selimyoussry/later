package boltdb_single

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hippoai/later/structures"
)

// CreateInstance creates a new instance in the database
func (database *Database) CreateInstance(name string, executionTime time.Time, parameters []byte) (string, error) {

	var instanceID string

	// Execute the transaction and store the instance ID outside the scope if successful
	err := database.DB.Update(func(tx *bolt.Tx) error {

		// Find the bucket
		b := tx.Bucket(bucket(name))

		// Make the key and value
		instanceID = MakeID(executionTime)
		value := &structures.Instance{
			ExecutionTime: executionTime.UTC(),
			ID:            instanceID,
			Parameters:    parameters,
			TaskName:      name,
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
