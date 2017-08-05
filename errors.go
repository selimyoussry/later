package later

import "github.com/hippoai/goerr"

func Err_TaskNameAlreadyTaken(taskName string) error {
	return goerr.New(ErrName_TaskNameAlreadyTaken, map[string]interface{}{
		"task_name": taskName,
	})
}

func Err_WrongToken(token string) error {
	return goerr.New(ErrName_WrongToken, map[string]interface{}{
		"token": token,
	})
}
