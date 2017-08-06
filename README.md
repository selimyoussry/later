# Execute your task at specified times, and add custom behavior


## Task

A `Task` is a [Golang interface](https://gobyexample.com/interfaces) implements the methods:
* `GetName() string` returns the task name.
* `OnCreate(t time.Time, parameters interface{}) error` creates an instance of this task that will run at time t, and is given parameters `parameters` in the form of a standard `interface{}`. In the task definition the proper Golang type of `parameters` should be specified so that the `OnCreate` function can do a type assertion and we're all happy. This parameters will be used to store all the necessary information about this task on runtime, and also for the callbacks. It should persist it to the database.
* `OnFail() error` defines the behavior when the task could not be run. By default it will print to log. It should persist it to the database.
* `OnSuccess() error` defines the behavior when the task was successfully implemented. It should persist it to the database.
* `OnAbort() error` defines the behavior when the task was aborted. It should persist it to the database.


## Instance

An instance is an instance of a specific task. It implements:
* `GetID() string` returns the unique ID automatically given to this instance upon creation.
* `GetTaskName() string` which returns the given task name.
* `GetExecutionTime() time.Time` which returns the execution time for this instance.
* `GetParameters() interface{}` which returns the parameters stored when the instance was created.


## Persistence

We store the tasks instances in a database. We don't want to restrict you to any database, we just require your database to implement the following methods:
* `GetInstances(start, end time.Time) ([]Instance, error)` which returns all the instances between a start and end time.
* `PullInstances(start, end time.Time) ([]Instance, error)` which returns all the instances between a start and end time, should mark them as pulled from the database, and puts them in the machine go-routines.
* `CreateInstance(name string, parameters interface{}) error` creates a task given passed parameters.
* `AbortInstance(name string, parameters interface{}) error` to abort a specific task instance.
* `GetLastPullTime() (time.Time, error)` returns the last time we pulled instances out of this database
* `SetPullTime(t time.Time) error` sets the pull to the given time.


# To Do

- [ ] Clean up code and isolate functionalities and pipelines
- [ ] Separate BoltDB database as standalone, open over gRPC
- [ ] Add BoltDB stats views and HTTP endpoint
- [ ] Secure the transactions
- [ ] Make it cluster-able
- [ ] Tutorial in the Readme
- [ ] Dockerize everything
- [ ] Add simple UI
