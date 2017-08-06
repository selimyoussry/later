package later

// AbortInstances removes the instances locally and from the database
func (machine *Machine) AbortInstance(name string, instanceID string) error {

	// Remove all these instances from the database
	err := machine.Database.AbortInstance(name, instanceID)
	if err != nil {
		return err
	}

	instanceItf, exists := machine.Instances.Load(instanceID)

	// If it exsists locally we close it gracefully and then delete
	// it from the list of instances
	if exists {

		// Abort the instance run
		localInstance := instanceItf.(*LocalInstance)
		localInstance.AbortLocally()

		// Delete it from the local list of instances
		machine.Instances.Delete(instanceID)
	}

	return nil
}
