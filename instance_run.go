package later

import (
	"time"

	"github.com/hippoai/goutil"
)

// Run runs the instance at the specific time unless aborted before
// and then closes the Abort channel
func (mi *ManagedInstance) Run() {
	defer close(mi.AbortChannel)

	wait := mi.Instance.ExecutionTime.Sub(time.Now())
	timer := time.NewTimer(wait)
	var err error

	select {

	// 1 - The task is due, execute it
	case <-timer.C:

		// Run the task
		err = mi.Task.Run()

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
			return
		}

		// Otherwise run OnSuccess
		err = mi.Task.OnSuccess()
		if err != nil {
			goutil.Log("Error on success %s - Got %s",
				mi.Instance.ID,
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
			goutil.Stringify(mi.Instance.Parameters),
		)

	}

}
