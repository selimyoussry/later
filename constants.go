package later

const (
	ErrName_TaskNameAlreadyTaken = "Err_TaskNameAlreadyTaken"
	ErrName_WrongToken           = "Err_WrongToken"

	Env_SecretKey = "SECRET_KEY"
	Env_Verbosity = "LATER_VERBOSE"

	gRPC_Server_Port = 9081
	HTTP_Server_Port = 8081

	STATUS_ABORTED    = "aborted"
	STATUS_FAILED     = "failed"
	STATUS_PENDING    = "pending"
	STATUS_SUCCESSFUL = "successful"

	DEFAULT_RECURRENCE_MIN = 5
	DEFAULT_TIME_AHEAD_MIN = 10
)
