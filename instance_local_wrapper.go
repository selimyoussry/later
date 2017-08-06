package later

import "github.com/hippoai/later/structures"

// LocalInstance is a wrapper around an instance that allows the Machine to store
// in one structure what's needed to run and stop an instance locally
type LocalInstance struct {
	AbortChannel chan bool
	Instance     *structures.Instance
	Task         Task
}

// NewLocalInstance instanciates a new local instance
// It essentially just creates the abort channel
func NewLocalInstance(instance *structures.Instance, task Task) *LocalInstance {
	abortChannel := make(chan bool)

	return &LocalInstance{
		AbortChannel: abortChannel,
		Instance:     instance,
		Task:         task,
	}
}

// Abort sends true to the abort channel, which causes the instance
// to stop and then close the AbortChannel
func (li *LocalInstance) AbortLocally() {
	li.AbortChannel <- true
}
