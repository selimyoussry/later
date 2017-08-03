package later

import (
	"time"
)

// UpdateLatestPullTime to be in sync with the database
// or to initialize it if it's the very first time we run this
func (machine *Machine) UpdateLatestPullTime() error {

	// See what's in the database
	latestPullTime, err := machine.Database.GetLastPullTime()
	if err != nil {
		return err
	}

	// If there was nothing in the database, update to now
	if latestPullTime == nil {
		now := time.Now()
		latestPullTime = &now
		machine.Database.SetPullTime(*latestPullTime)
	}

	machine.LatestPullTime = latestPullTime
	return nil

}

// LoopCore is the core of the loop
func (machine *Machine) LoopCore() error {

	// 1 - We get all the pending instances for the current timeframe
	instances, err := machine.Database.GetInstances(
		machine.LatestPullTime.Add(-1*time.Minute),
		machine.LatestPullTime.Add(machine.Parameters.TimeAhead),
	)
	if err != nil {
		return err
	}

	// 2 - We add the new instances
	for _, instance := range instances {
		machine.StartInstance(instance)
	}

	// 3 - We wait until next fetch
	timeOfNextRun := machine.LatestPullTime.Add(machine.Parameters.Recurrence)
	wait := timeOfNextRun.Sub(time.Now())
	time.Sleep(wait)

	// We set the new pull time in the database and call this function again
	now := time.Now()
	machine.Database.SetPullTime(now)
	machine.LatestPullTime = &now

	return nil

}

// Loop infinitely adds and runs new instances
func (machine *Machine) Loop() error {

	err := machine.UpdateLatestPullTime()
	if err != nil {
		return err
	}

	for {
		err := machine.LoopCore()
		if err != nil {
			return err
		}
	}

}
