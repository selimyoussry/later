package later

// RegisterTasks registers all the declared tasks
func (machine *Machine) RegisterTasks(tasks ...Task) error {

	// Loop over the tasks and register them
	for _, task := range tasks {

		err := machine.registerTask(task)
		if err != nil {
			return err
		}

	}

	return nil
}

// registerTask registers a single task in the machine
func (machine *Machine) registerTask(task Task) error {
	// Check the task name has not already been used
	taskName := task.GetName()
	_, exists := machine.Tasks[taskName]
	if exists {
		return Err_TaskNameAlreadyTaken(taskName)
	}

	// Save it in the tasks map
	machine.Tasks[taskName] = task
	return nil

}
