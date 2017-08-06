package boltdb_app_server

import "fmt"

const (
	Env_Path     = "BOLTDB_PATH"
	Default_Path = "./"
	DB_File_Name = "later.bolt.db"

	BUCKET_ABORTED    = "__Aborted"
	BUCKET_SUCCESSFUL = "__Successful"
	BUCKET_FAILED     = "__Failed"
	BUCKET_METADATA   = "__Metadata"

	KEY_LAST_PULL_TIME = "LastPullTime"

	gRPC_Server_Port = 9080
	HTTP_Server_Port = 8080
)

func bucket(taskName string) []byte {
	return []byte(fmt.Sprintf("Instances.%s", taskName))
}
