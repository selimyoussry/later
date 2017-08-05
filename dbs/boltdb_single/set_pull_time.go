package boltdb_single

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

// SetPullTime
func (database *Database) SetPullTime(t time.Time) error {

	err := database.DB.Update(func(tx *bolt.Tx) error {

		metadataBucket := tx.Bucket([]byte(BUCKET_METADATA))

		b, err := json.Marshal(&t)
		if err != nil {
			return err
		}

		err = metadataBucket.Put([]byte(KEY_LAST_PULL_TIME), b)
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
