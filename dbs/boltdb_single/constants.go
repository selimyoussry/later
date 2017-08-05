package boltdb_single

import "fmt"

const (
	BUCKET_ABORTED    = "__Aborted"
	BUCKET_SUCCESSFUL = "__Successful"
	BUCKET_FAILED     = "__Failed"
	BUCKET_METADATA   = "__Metadata"

	KEY_LAST_PULL_TIME = "LastPullTime"
)

func bucket(taskName string) []byte {
	return []byte(fmt.Sprintf("Instances.%s", taskName))
}
