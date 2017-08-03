package later

import "time"

type Instance struct {
	ExecutionTime time.Time
	ID            string
	Parameters    interface{}
	TaskName      string
}

type ManagedInstance struct {
	AbortChannel chan bool
	Instance     *Instance
	Task         Task
}

// NewManagedInstance instanciates a new managed instance
func NewManagedInstance(instance *Instance, task Task) *ManagedInstance {
	abortChannel := make(chan bool)

	return &ManagedInstance{
		AbortChannel: abortChannel,
		Instance:     instance,
		Task:         task,
	}
}

// Abort sends true to the abort channel, which causes the instance
// to stop and then close the AbortChannel
func (mi *ManagedInstance) Abort() {
	mi.AbortChannel <- true
}

// StartInstance stores and starts a particular instance
func (machine *Machine) StartInstance(instance *Instance) {
	task := machine.Tasks[instance.TaskName]
	managedInstance := NewManagedInstance(instance, task)
	machine.Instances.Store(instance.ID, managedInstance)

	// Launch a goroutine for this instance - which will remove itself
	// from the map when run
	go func() {
		defer machine.Instances.Delete(instance.ID)
		managedInstance.Run()
	}()
}

// AbortInstance removes the instances locally and from the database
func (machine *Machine) AbortInstance(name string, parameters interface{}) error {

	// Remove from the database
	instanceID, err := machine.Database.AbortInstance(name, parameters)
	if err != nil {
		return err
	}

	// Remove locally
	instanceItf, exists := machine.Instances.Load(instanceID)

	// If it exsists locally we close it gracefully and then delete
	// it from the list of instances
	if exists {
		managedInstance := instanceItf.(*ManagedInstance)
		managedInstance.Abort()
		machine.Instances.Delete(instanceID)
		machine.Instances.Delete(instanceID)
	}

	return nil
}
