package structures

import "time"

type Instance struct {
	ExecutionTime time.Time
	ID            string
	Parameters    []byte
	TaskName      string
}
