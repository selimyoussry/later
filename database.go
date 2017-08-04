package later

import "time"

// Database lists what the database needs to be able to do for this library
type Database interface {
	AbortInstances(name string, parameters []byte) ([]string, error)
	CreateInstance(name string, executionTime time.Time, parameters []byte) (string, error)
	GetInstances(start, end time.Time) ([]*Instance, error)
	GetLastPullTime() (*time.Time, error)
	SetPullTime(t time.Time) error
}
