package later

import "time"

type Task interface {
	GetName() string
	GetExecutionTime() time.Time
	GetParameters() interface{}

	OnCreate(executionTime time.Time, parameters interface{}) error
	OnFail(runError error) error
	OnSuccess() error
	OnAbort() error

	Run() error
}
