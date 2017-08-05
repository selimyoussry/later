package later

// GetNumberOfLocalInstances returns the number of instances
func (machine *Machine) GetNumberOfLocalInstances() int64 {

	size := int64(0)
	machine.Instances.Range(func(key interface{}, value interface{}) bool {
		size += 1
		return true
	})

	return size

}
