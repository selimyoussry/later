package later

import "time"

type Instance interface {
	GetExecutionTime() time.Time
	GetID() string
	GetParameters() []byte
	GetTaskName() string
}

type ManagedInstance struct {
	AbortChannel chan bool
	Instance     Instance
	Task         Task
}

// NewManagedInstance instanciates a new managed instance
func NewManagedInstance(instance Instance, task Task) *ManagedInstance {
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
func (machine *Machine) StartInstance(instance Instance) {
	task := machine.Tasks[instance.GetTaskName()]
	managedInstance := NewManagedInstance(instance, task)
	machine.Instances.Store(instance.GetID(), managedInstance)

	// Launch a goroutine for this instance - which will remove itself
	// from the map when run
	go func() {
		defer machine.Instances.Delete(instance.GetID())
		managedInstance.Run()
	}()
}

// AbortInstance removes the instances locally and from the database
func (machine *Machine) AbortInstances(name string, parameters []byte) ([]string, error) {

	// Remove from the database
	instancesIDs, err := machine.Database.AbortInstances(name, parameters)
	if err != nil {
		return []string{}, err
	}

	// Remove locally each instance
	for _, instanceID := range instancesIDs {

		instanceItf, exists := machine.Instances.Load(instanceID)

		// If it exsists locally we close it gracefully and then delete
		// it from the list of instances
		if exists {
			managedInstance := instanceItf.(*ManagedInstance)
			managedInstance.Abort()
			machine.Instances.Delete(instanceID)
			machine.Instances.Delete(instanceID)
		}

	}

	return instancesIDs, nil
}
