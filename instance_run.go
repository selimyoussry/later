package later

import (
	"time"

	"github.com/hippoai/goutil"
	"github.com/hippoai/later/structures"
)

// RunInstance runs an instance locally
func (machine *Machine) RunInstanceIfNotAlreadyThere(instance *structures.Instance) {

	// Check if the instance is already in the local memory
	_, exists := machine.Instances.Load(instance.ID)
	if exists {
		return
	}

	// Wrap the instance in a "local instance" that has an abort channel
	task := machine.Tasks[instance.TaskName]
	localInstance := NewLocalInstance(instance, task)
	machine.Instances.Store(instance.ID, localInstance)

	// At the end of the run, close the channel and delete the instance it from the local store
	defer func() {
		close(localInstance.AbortChannel)
		machine.Instances.Delete(instance.ID)
	}()

	// Create a timer to trigger the instance at the right time
	wait := instance.ExecutionTime.Sub(time.Now())
	timer := time.NewTimer(wait)

	// Execute or Abort the instance
	select {

	// Execute the instance
	case <-timer.C:

		// Run the task
		response, err := task.Run(instance.Parameters)

		// If there is an error on run
		if err != nil {

			// Log the error
			goutil.Log("Error on running instance %s - Got %s",
				instance.ID,
				goutil.Stringify(err),
			)

			// Run OnFail callback for this task
			err = task.OnFail(err)
			if err != nil {
				goutil.Log("Error on failing %s - Got %s",
					instance.ID,
					goutil.Stringify(err),
				)
			}

			// Save the failed instance in the database
			if err != nil {
				goutil.Log("Error on saving failed to db %s",
					goutil.Stringify(err),
				)
			}

			// Exit the function
			return

		}

		// No error on run, this instance has successfully completed

		// Run OnSuccess callback
		err = task.OnSuccess(response)
		if err != nil {
			goutil.Log("Error on success %s - Got %s",
				instance.ID,
				goutil.Stringify(err),
			)
		}

		// Save the success in the database
		err = machine.Database.MarkAsSuccessful(task.GetName(), instance.ID)
		if err != nil {
			goutil.Log("Error on saving success to db %s",
				goutil.Stringify(err),
			)
		}

		// Abort this instance
		// This can only happen if the "AbortInstance" API has been called
		// This just prevents the instance from running locally
		// It is removed from the database by the "AbortInstance" function
	case <-localInstance.AbortChannel:
		goutil.Log("Aborting instance %s at %s | Task %s scheduled for %s with parameters %s",
			instance.ID,
			time.Now().String(),
			task.GetName(),
			instance.ExecutionTime.String(),
			string(instance.Parameters),
		)

	}

}
