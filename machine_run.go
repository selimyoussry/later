package later

import (
	"time"

	"github.com/hippoai/later/structures"
)

// Loop infinitely adds and runs new instances
func (machine *Machine) Loop() error {

	// First, update the latest information from the database
	err := machine.UpdateLatestPullTime()
	if err != nil {
		return err
	}

	// Then loop
	for {
		err := machine.LoopCore()
		if err != nil {
			return err
		}
	}

}

// LoopCore is the core of the loop
func (machine *Machine) LoopCore() error {

	// 1 - We get all the pending instances for the current timeframe
	instances, err := machine.GetInstances(
		machine.LatestPullTime.Add(-1*time.Minute),
		machine.LatestPullTime.Add(machine.Parameters.TimeAhead),
	)
	if err != nil {
		return err
	}

	// 2 - We add the new instances
	// spawning a new go-routine for each of them
	for _, instance := range instances {
		go func(instance *structures.Instance) {
			machine.RunInstanceIfNotAlreadyThere(instance)
		}(instance)
	}

	// 3 - We wait until next fetch
	timeOfNextRun := machine.LatestPullTime.Add(machine.Parameters.Recurrence)
	wait := timeOfNextRun.Sub(time.Now())
	time.Sleep(wait)

	// We set the new pull time in the database and call this function again
	now := time.Now()
	err = machine.Database.SetPullTime(now)
	if err != nil {
		return err
	}

	machine.LatestPullTime = &now

	return nil

}
