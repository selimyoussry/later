package later

import (
	"time"

	"github.com/hippoai/goutil"
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
		err = machine.Database.SetPullTime(*latestPullTime)
		if err != nil {
			goutil.Log("Could not set pull time to %s",
				goutil.Stringify(now),
			)
			return err
		}
	}

	machine.LatestPullTime = latestPullTime
	return nil

}
