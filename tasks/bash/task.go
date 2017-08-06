package bash

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/hippoai/goutil"
	"github.com/hippoai/later/tasks"
)

type Task struct {
}

type Parameters struct {
	Command           *string `json:"command"`
	OnFailEndpoint    *string `json:"onFailEndpoint"`
	OnSuccessEndpoint *string `json:"onSuccessEndpoint"`
	OnAbortEndpoint   *string `json:"onAbortEndpoint"`
}

type Response struct {
	Stdout []byte `json:"stdout"`
	Stderr []byte `json:"stderr"`
}

func (task *Task) GetName() string {
	return "bash"
}

func (task *Task) OnFail(runError error) error {

	return nil
}

func (task *Task) OnSuccess(responseItf interface{}) error {

	response := responseItf.(*Response)
	log.Printf("Task: Bash - Got response %s \n", goutil.Pretty(response))

	return nil
}

func (task *Task) OnAbort() error {

	return nil
}

// Run -
func (task *Task) Run(parametersAsBytes []byte) (interface{}, error) {

	var parameters Parameters
	err := json.Unmarshal(parametersAsBytes, &parameters)
	if err != nil {
		return nil, err
	}
	if parameters.Command == nil {
		return nil, tasks.Err_PayloadDecode("task: bash")
	}

	// Execute the bash command
	cmd := exec.Command("bash", "-c", *parameters.Command)
	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// Get stdout and stderr while it's running
	stdout, err := ioutil.ReadAll(stdoutReader)
	if err != nil {
		return nil, err
	}

	stderr, err := ioutil.ReadAll(stderrReader)
	if err != nil {
		return nil, err
	}

	// Get running error code if any
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return &Response{
		Stdout: stdout,
		Stderr: stderr,
	}, nil
}
