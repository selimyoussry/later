package boltdb_app_server

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

// GetLastPullTime
func (database *Database) GetLastPullTime() (*time.Time, error) {

	var lastPullTime time.Time
	defined := false

	err := database.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(BUCKET_METADATA))
		lastPullTimeAsBytes := b.Get([]byte(KEY_LAST_PULL_TIME))

		// We never set this time
		if lastPullTimeAsBytes == nil {
			return nil
		}

		// Otherwise we set it, so we need to unmarshal it
		err := json.Unmarshal(lastPullTimeAsBytes, &lastPullTime)
		if err != nil {
			return err
		}

		defined = true
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Not every defined, return nil
	if !defined {
		return nil, nil
	}

	return &lastPullTime, nil
}
