package boltdb_app_server

import (
	"path"

	"github.com/hippoai/env"
)

func GetPath() string {

	parsed, err := env.Parse(Env_Path)
	if err != nil {
		return path.Join(Default_Path, DB_File_Name)
	}

	return path.Join(parsed[Env_Path], DB_File_Name)

}
