package echo

import (
	"encoding/json"
	"log"

	"github.com/hippoai/later/tasks"
)

type Task struct {
}

type Parameters struct {
	Message *string `json:"message"`
}

func (task *Task) GetName() string {
	return "echo"
}

func (task *Task) OnFail(runError error) error {
	log.Println("Echo failed")
	return nil
}

func (task *Task) OnSuccess(response interface{}) error {
	log.Println("Echo succeeded")

	return nil
}

func (task *Task) OnAbort() error {
	log.Println("Echo aborted")
	return nil
}

func (task *Task) Run(parametersAsBytes []byte) (interface{}, error) {

	var parameters Parameters
	err := json.Unmarshal(parametersAsBytes, &parameters)
	if err != nil {
		return nil, tasks.Err_PayloadDecode("task: echo")
	}

	log.Printf("Running echo: %s \n", *parameters.Message)
	return nil, nil
}
