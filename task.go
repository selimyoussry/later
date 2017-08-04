package later

type Task interface {
	GetName() string

	OnFail(runError error) error
	OnSuccess(response interface{}) error
	OnAbort() error

	Run(parametersAsBytes []byte) (interface{}, error)
}
