package boltdb

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/hippoai/goutil"
)

// moveInstances from one bucket to another
func (database *Database) moveInstances(srcBucketName, dstBucketName []byte, instancesIDs []string) ([]string, error) {

	errorsDstBucket := []string{}

	database.DB.Update(func(tx *bolt.Tx) error {

		var err error

		srcBucket := tx.Bucket(srcBucketName)
		dstBucket := tx.Bucket(dstBucketName)

		// For each instance ID, delete it and place it in a bucket
		for _, instanceID := range instancesIDs {

			// Get the value
			key := []byte(instanceID)
			valueAsBytes := srcBucket.Get(key)

			// Set the value in aborted
			err = dstBucket.Put(key, valueAsBytes)
			if err != nil {
				errorsDstBucket = append(errorsDstBucket, instanceID)
			}

			// Delete from list of tasks
			err = srcBucket.Delete(key)
			if err != nil {
				return err
			}

		}

		return nil

	})

	if len(errorsDstBucket) > 0 {
		log.Printf("Could not put in destination bucket %s: %s \n", string(dstBucketName), goutil.Stringify(errorsDstBucket))
	}

	return instancesIDs, nil

}
