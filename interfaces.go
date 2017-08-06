package later

import (
	"time"

	"github.com/hippoai/later/structures"
)

// Database lists what the database needs to be able to do for this library
// note that the database might have extra functionality, these are minimum compatibility requirements
type Database interface {
	AbortInstance(taskName string, instanceID string) error
	Close() error
	CreateInstance(taskName string, executionTime time.Time, parameters []byte) (string, error)
	GetInstances(start, end time.Time) ([]*structures.Instance, error)
	GetLastPullTime() (*time.Time, error)
	MarkAsSuccessful(taskName string, instanceID string) error
	MarkAsFailed(taskName string, instanceID string) error
	SetPullTime(t time.Time) error
}

// Task defines what be registered when creating this job scheduler
// This comes packages with "echo" (for testing) and "bash" tasks
type Task interface {
	GetName() string

	OnFail(runError error) error
	OnSuccess(response interface{}) error
	OnAbort() error

	Run(parametersAsBytes []byte) (interface{}, error)
}
