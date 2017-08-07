package boltdb_app_server

const (
	Env_Path     = "BOLTDB_PATH"
	Default_Path = "./data"
	DB_File_Name = "later.bolt.db"

	BUCKET_ABORTED    = "__Aborted"
	BUCKET_SUCCESSFUL = "__Successful"
	BUCKET_FAILED     = "__Failed"
	BUCKET_PENDING    = "__Pending"

	KEY_LAST_PULL_TIME = "LastPullTime"

	gRPC_Server_Port      = 9080
	HTTP_Server_Port      = 8080
	HTTP_gRPC_Server_Port = 8081
)
