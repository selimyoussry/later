package later

import (
	"time"

	"github.com/hippoai/later/structures"
)

func (machine *Machine) GetInstances(start, end time.Time) ([]*structures.Instance, error) {

	// Get the instances from the database during this timeframe
	instances, err := machine.Database.GetInstances(start, end)
	if err != nil {
		return nil, err
	}

	return instances, nil

}
