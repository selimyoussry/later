package later

type Task interface {
	GetName() string

	OnFail(runError error) error
	OnSuccess() error
	OnAbort() error

	Run(parametersItf interface{}) error
}
