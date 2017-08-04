package boltdb

import "fmt"

const (
	BUCKET_ABORTED   = "__Aborted"
	BUCKET_COMPLETED = "__Completed"
	BUCKET_METADATA  = "__Metadata"

	KEY_LAST_PULL_TIME = "LastPullTime"
)

func bucket(taskName string) []byte {
	return []byte(fmt.Sprintf("Instances.%s", taskName))
}
