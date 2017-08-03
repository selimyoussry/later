package later

import "time"

// Database lists what the database needs to be able to do for this library
type Database interface {
	AbortInstance(name string, parameters interface{}) (string, error)
	CreateInstance(name string, parameters interface{}) (string, error)
	GetInstances(start, end time.Time) ([]*Instance, error)
	GetLastPullTime() (*time.Time, error)
	SetPullTime(t time.Time) error
}
