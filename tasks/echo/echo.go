package echo

import (
	"log"

	"github.com/hippoai/goerr"
)

type Task struct {
}

type Parameters struct {
	Message string `json:"message"`
}

func (task *Task) GetName() string {
	return "echo"
}

func (task *Task) OnFail(runError error) error {
	log.Println("Echo failed")
	return nil
}

func (task *Task) OnSuccess() error {
	log.Println("Echo succeeded")

	return nil
}

func (task *Task) OnAbort() error {
	log.Println("Echo aborted")
	return nil
}

func (task *Task) Run(parametersItf interface{}) error {

	parameters, ok := parametersItf.(*Parameters)
	if !ok {
		return goerr.New("ERR_CAST", map[string]interface{}{
			"task": "echo",
		})
	}

	log.Printf("Running echo: %s \n", parameters.Message)

	return nil
}
