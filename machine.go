package later

import (
	"time"

	"golang.org/x/sync/syncmap"
)

// Machine is the program that initializes the tasks, connects to the database
// fetches the instances, and makes sure everything is ran properly
type Machine struct {
	Database  Database    // The database we plugged in
	Instances syncmap.Map // All the instances waiting to be executed in our go-routines

	LatestPullTime *time.Time
	Parameters     *MachineParameters
	Tasks          map[string]Task // All the tasks we've declared at compile time
}

// MachineParameters defines how often we pull new tasks
type MachineParameters struct {
	Recurrence time.Duration // We're going to pull tasks every PullRecurrence
	TimeAhead  time.Duration //  We're going to pull tasks for the next TimeAhead
}

func GetDefaultMachineParameters() *MachineParameters {
	return &MachineParameters{
		Recurrence: 1 * time.Minute,
		TimeAhead:  2 * time.Minute,
	}
}

// NewMachine instanciates
func NewMachine(database Database, parameters *MachineParameters) *Machine {

	if parameters == nil {
		parameters = GetDefaultMachineParameters()
	}

	return &Machine{
		Database:       database,
		Instances:      syncmap.Map{},
		LatestPullTime: nil,
		Tasks:          map[string]Task{},
		Parameters:     parameters,
	}
}
