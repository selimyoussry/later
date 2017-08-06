package later

import "time"

// MachineParameters defines how often we pull new tasks
type MachineParameters struct {
	Recurrence time.Duration // We're going to pull tasks every PullRecurrence
	TimeAhead  time.Duration //  We're going to pull tasks for the next TimeAhead
}

// GetDefaultMachineParameters creates default machine parameters
func GetDefaultMachineParameters() *MachineParameters {
	return &MachineParameters{
		Recurrence: 1 * time.Minute,
		TimeAhead:  2 * time.Minute,
	}
}
