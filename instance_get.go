package later

import (
	"time"

	"github.com/hippoai/later/structures"
)

func (machine *Machine) getter(start, end time.Time, status string) ([]*structures.Instance, error) {

	switch status {

	case STATUS_FAILED:
		return machine.Database.GetFailed(start, end)

	case STATUS_SUCCESSFUL:
		return machine.Database.GetSuccessful(start, end)

	case STATUS_ABORTED:
		return machine.Database.GetAborted(start, end)

	}

	return machine.Database.GetInstances(start, end)

}
