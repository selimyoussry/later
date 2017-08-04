package tasks

import "github.com/hippoai/goerr"

func Err_PayloadDecode(endpoint string) error {

	return goerr.New(ErrName_PayloadDecode, map[string]interface{}{
		"endpoint": endpoint,
	})

}
