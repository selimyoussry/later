package boltdb_app_server

import (
	"path"

	"github.com/hippoai/env"
)

func GetPath() string {
	parsed, err := env.Parse(Env_Path)
	if err != nil {
		return Default_Path
	}

	return parsed[Env_Path]
}

func GetFilePath() string {
	return path.Join(GetPath(), DB_File_Name)
}
