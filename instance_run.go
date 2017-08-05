package later

import (
	"time"

	"github.com/hippoai/goutil"
)

// Run runs the instance at the specific time unless aborted before
// and then closes the Abort channel
func (mi *ManagedInstance) Run(machine *Machine) {
	defer close(mi.AbortChannel)

	wait := mi.Instance.ExecutionTime.Sub(time.Now())
	timer := time.NewTimer(wait)

	select {

	// 1 - The task is due, execute it
	case <-timer.C:

		// Run the task
		response, err := mi.Task.Run(mi.Instance.Parameters)

		// If there is an error on run, call the OnFail callback
		if err != nil {
			goutil.Log("Error on running instance %s - Got %s",
				mi.Instance.ID,
				goutil.Stringify(err),
			)

			// Run OnFail
			err = mi.Task.OnFail(err)
			if err != nil {
				goutil.Log("Error on failing %s - Got %s",
					mi.Instance.ID,
					goutil.Stringify(err),
				)
			}

			// Save the failed in the database
			err = machine.Database.MarkAsFailed(mi.Task.GetName(), mi.Instance.ID)
			if err != nil {
				goutil.Log("Error on saving failed to db %s",
					goutil.Stringify(err),
				)
			}

			return
		}

		// Otherwise run OnSuccess
		err = mi.Task.OnSuccess(response)
		if err != nil {
			goutil.Log("Error on success %s - Got %s",
				mi.Instance.ID,
				goutil.Stringify(err),
			)
		}

		// Save the success in the database
		err = machine.Database.MarkAsSuccessful(mi.Task.GetName(), mi.Instance.ID)
		if err != nil {
			goutil.Log("Error on saving success to db %s",
				goutil.Stringify(err),
			)
		}

		// 2 - It was aborted
	case <-mi.AbortChannel:
		goutil.Log("Aborting instance %s at %s | Task %s scheduled for %s with parameters %s",
			mi.Instance.ID,
			time.Now().String(),
			mi.Task.GetName(),
			mi.Instance.ExecutionTime.String(),
			string(mi.Instance.Parameters),
		)

	}

}
