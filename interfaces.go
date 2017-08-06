package later

import (
	"time"

	"github.com/hippoai/later/structures"
)

// Database lists what the database needs to be able to do for this library
// note that the database might have extra functionality, these are minimum compatibility requirements
type Database interface {
	AbortInstance(instanceID string) error
	Close() error
	CreateInstance(taskName string, executionTime time.Time, parameters []byte) (string, error)
	GetInstances(start, end time.Time) ([]*structures.Instance, error)
	GetAborted(start, end time.Time) ([]*structures.Instance, error)
	GetSuccessful(start, end time.Time) ([]*structures.Instance, error)
	GetFailed(start, end time.Time) ([]*structures.Instance, error)
	MarkAsSuccessful(instanceID string) error
	MarkAsFailed(instanceID string) error
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
