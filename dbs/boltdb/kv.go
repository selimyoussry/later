package boltdb

import (
	"fmt"
	"time"

	"github.com/hippoai/goutil"
)

// MakeID creates a unique ID, sortable by execution time
func MakeID(executionTime time.Time) string {
	uuid := goutil.UuidV4()

	return fmt.Sprintf("%s.%s",
		executionTime.Format(time.RFC3339),
		uuid,
	)
}

type Value struct {
	ExecutionTime string `json:"executionTime"`
	ID            string `json:"id"`
	Parameters    []byte `json:"parameters"`
}
