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
