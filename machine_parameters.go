package later

import "time"

// MachineParameters defines how often we pull new tasks
type machineParameters struct {
	Recurrence time.Duration // We're going to pull tasks every PullRecurrence
	TimeAhead  time.Duration //  We're going to pull tasks for the next TimeAhead
}

func NewMachineParameters(recurrence, timeAhead time.Duration) *machineParameters {
	return &machineParameters{
		Recurrence: recurrence,
		TimeAhead:  timeAhead,
	}
}

// GetDefaultMachineParameters creates default machine parameters
func GetDefaultMachineParameters() *machineParameters {
	return &machineParameters{
		Recurrence: DEFAULT_RECURRENCE_MIN * time.Minute,
		TimeAhead:  DEFAULT_TIME_AHEAD_MIN * time.Minute,
	}
}
