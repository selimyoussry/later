package later

import "github.com/hippoai/goerr"

func Err_TaskNameAlreadyTaken(taskName string) error {
	return goerr.New(ErrName_TaskNameAlreadyTaken, map[string]interface{}{
		"task_name": taskName,
	})
}
