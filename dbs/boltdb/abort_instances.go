package boltdb

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
	"github.com/hippoai/goutil"
)

type Input struct {
	InstancesIDs []string `json:"instancesIDs"`
}

// AbortInstances should be able to abort the instances with either
// - a list of IDs
// - all for a given task
// - between a start and end time
func (database *Database) AbortInstances(taskName string, parameters []byte) ([]string, error) {

	var input Input
	err := json.Unmarshal(parameters, &input)
	if err != nil {
		return []string{}, err
	}

	errorsAbortedBucket := []string{}

	database.DB.Update(func(tx *bolt.Tx) error {

		var err error

		taskBucket := tx.Bucket(bucket(taskName))
		abortedBucket := tx.Bucket(bucket(BUCKET_ABORTED))

		// For each instance ID, delete it and place it in a bucket
		for _, instanceID := range input.InstancesIDs {

			// Get the value
			key := []byte(instanceID)
			valueAsBytes := taskBucket.Get(key)

			// Set the value in aborted
			err = abortedBucket.Put(key, valueAsBytes)
			if err != nil {
				errorsAbortedBucket = append(errorsAbortedBucket, instanceID)
			}

			// Delete from list of tasks
			err = taskBucket.Delete(key)
			if err != nil {
				return err
			}

		}

		return nil

	})

	if len(errorsAbortedBucket) > 0 {
		log.Printf("Could not put in aborted bucket: %s \n", goutil.Stringify(errorsAbortedBucket))
	}

	return input.InstancesIDs, nil

}
