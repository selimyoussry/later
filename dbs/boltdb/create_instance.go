package boltdb

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
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

	return instanceID, err

}
