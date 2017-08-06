package boltdb_app_server

import (
	"fmt"
	"time"

	"github.com/hippoai/goutil"
)

// MakeID creates a unique ID, sortable by execution time
func MakeID(executionTime time.Time, taskName string) string {
	uuid := goutil.UuidV4()

	return fmt.Sprintf("%s.%s.%s",
		executionTime.UTC().Format(time.RFC3339),
		taskName,
		uuid,
	)
}
