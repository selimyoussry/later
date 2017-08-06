package boltdb_app_server

import "github.com/boltdb/bolt"

func Initialize(db *bolt.DB) error {

	buckets := []string{
		BUCKET_ABORTED,
		BUCKET_FAILED,
		BUCKET_PENDING,
		BUCKET_SUCCESSFUL,
	}

	// Initialize the buckets
	err := db.Update(func(tx *bolt.Tx) error {

		var err error

		// Create the buckets
		for _, bucketName := range buckets {

			// Create the aborted bucket
			_, err = tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}

		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil

}
