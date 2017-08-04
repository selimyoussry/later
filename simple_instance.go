package later

import "time"

type SimpleInstance struct {
	ExecutionTime time.Time
	ID            string
	Parameters    []byte
	TaskName      string
}

func (si *SimpleInstance) GetExecutionTime() time.Time {
	return si.ExecutionTime
}

func (si *SimpleInstance) GetID() string {
	return si.ID
}

func (si *SimpleInstance) GetParameters() []byte {
	return si.Parameters
}

func (si *SimpleInstance) GetTaskName() string {
	return si.TaskName
}
