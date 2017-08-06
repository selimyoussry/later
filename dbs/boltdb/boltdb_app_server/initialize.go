package boltdb_app_server

import "github.com/boltdb/bolt"

func Initialize(db *bolt.DB, tasks []string) error {

	// Initialize the buckets
	err := db.Update(func(tx *bolt.Tx) error {

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
		return err
	}

	return nil

}
