# Later

Later is a job scheduler, aiming at becoming a simple replacement for Cron. It uses go-routines to run tasks concurrently and only pulls in memory jobs that will run in the following hours or so (this is configurable), so it should be able to scale to millions of scheduled jobs with a single node.


## Project status
It works on a single node right now but will eventually be distributed. It is still a work in progress, but should be in a working state as of now. Please submit issues as you find bugs. Please also note this was built as a job scheduler, not as a job runner, i.e. don't try to run long-running processes with this. Rather, create an API with your workers, and trigger workers using this library. This way, you separate job scheduling (with Later) from resource management for all your workers pool.


## Table of Contents

* [Getting Started](#getting-started)
  * [Installing](#installing)
  * [Write a Task](#write-a-task)
  * [Run it](#run-it)
    * Docker
  * [Database](#database)
    * Built-in BoltDB
    * Connect your own database
* [To Do](#to-do)


## Getting started

### Installing

To start using Later, install Go and run `go get`:
`$ go get github.com/hippoai/later`.

This will pull the job scheduler library. In order to persist jobs, we also need to use a database that implements the `Database` interface. You can use the built-in [BoltDB](https://github.com/boltdb/bolt) driver, or [build your own driver](#build-your-own-database-driver). See [Database](#database) section for more information on how to run a compatible database.


### Write a task

The two main concepts of this library are a `task` and an `instance`. A task is a job definition, it's got a name, what to do when it's called, what to do if it fails, succeeds, or is aborted. An instance is an instance of a task. For example, this library ships with the "echo" task, its name is "echo", it does nothing after it just prints "failed" when it fails, "success" when it succeeds or "aborted" when it is aborted. And when an instance of this task runs, it prints back whatever message was saved in its instance.

A task is a structure that implements the interface

```go
type Task interface {
	GetName() string

	OnFail(runError error) error
	OnSuccess(response interface{}) error
	OnAbort() error

	Run(parametersAsBytes []byte) (interface{}, error)
}
```

Write your own task by creating a structure that implements all these functions. This will generate a new job definition, and you will therefore be able to launch instances of this task that will execute at a given time. An instance has:
* a task name
* parameters (passed as a byte array, usually the serialization of some JSON)
* an execution time
* an ID (automatically generated upon creation)

Look at `/instances/bash` to see how running "bash" commands is implemented here.


### Run it

Once the code for your tasks is written, and you've picked a database driver, we are ready to run Later. In a `main` package, we first import the following (assuming you are using the BoltDB driver - just replace it with your own driver otherwise).

```go
"github.com/hippoai/later"
"github.com/hippoai/later/dbs/boltdb"
```

We are going to add the libraries for the tasks this ships with. You would add the paths to your own tasks here.
```go
"github.com/hippoai/later/tasks/bash"
"github.com/hippoai/later/tasks/echo"
```

```go
func main(){

  // Database driver
  db, err := boltdb.NewDatabaseFromEnv()
  if err != nil {
    log.Fatal(err)
  }

  // Create a machine with default parameters (nil)
  machine := later.NewMachine(db, nil)

  // Register the tasks you want to add, here we add the default bash and echo for the example
  err = machine.RegisterTasks(
    &bash.Task{},
    &echo.Task{},
  )
  if err != nil {
    log.Fatal(err)
  }

  // Then we listen for incoming connections on gRPC, port 9081
	gRPC_server := later.NewServer(machine, "")
	go func() {
		for {
			err := gRPC_server.Run_gRPC()
      log.Printf("Err with the gRPC server %s \n", err)
		}
	}()

  // We also serve on HTTP, port 8081
  go func() {
    for {
      err := gRPC_server.Run_HTTP()
      log.Printf("Error with HTTP server %s \n", err)
    }
  }()

  // And run the program - this is the job coordinator and executor
  err = machine.Loop()
  if err != nil {
    log.Fatal(err)
  }

}
```

#### Docker

If you'd like to Dockerize this program, you should expose ports 9081 (gRPC server) and 8081 (HTTP server).


#### Custom parameters

Later pulls the jobs for the next X minutes, every Y minutes. X is set by default to 10 minutes, and Y to 5 minutes. Depending on how much memory you have and how many jobs you are planning, you can customize these parameters by using the `NewMachineParameters(recurrence, timeAhead time.Duration)` instead of `nil` when creating a new machine with `later.NewMachine`.

Also, you can make this library not verbose with having an environment variable `LATER_VERBOSE=0`. By default it will log everything.

### Database

Later needs a database implementing the following interface to schedule jobs. We provide an application server using BoltDB that implements it, or you can write your own using your database of choice.


#### Built-in BoltDB

You can pull the Docker image `hippoai/later-boltdb-app-server`. It exposes a gRPC port on 9080 for communication with Later, and an HTTP port on 8080 that provides a `GET /export` endpoint to backup the database. Also, you should mount a volume on `/app/data` to store the database data outside of the Docker container. This looks like:
```
docker run -d -p 8080:8080 -p 9080:9080 -v ~/bolt:/app/data hippoai/later-boltdb-app-server
```

You can run database exports by calling `GET /export` method on port 8080, this will return a file. [See this](https://github.com/boltdb/bolt/blob/master/README.md#database-backups) for more info. This is a hot backup, which means your database won't be down during the time of the backup.


#### Connect your own database

Following the built-in BoltDB driver in `/dbs/boltdb`, your database driver needs to implement the following interface to be compatible:

```go
type Database interface {
	AbortInstance(instanceID string) error
	Close() error
	CreateInstance(taskName string, executionTime time.Time, parameters []byte) (string, error)
	GetInstances(start, end time.Time) ([]*structures.Instance, error)
	GetAborted(start, end time.Time) ([]*structures.Instance, error)
	GetSuccessful(start, end time.Time) ([]*structures.Instance, error)
	GetFailed(start, end time.Time) ([]*structures.Instance, error)
	MarkAsSuccessful(instanceID string) error
	MarkAsFailed(instanceID string) error
}
```


## To Do

- [x] Clean up code and isolate functionalities and pipelines
- [x] Separate BoltDB database as standalone, open over gRPC
- [ ] Secure the transactions
- [ ] Make it cluster-able
- [x] Tutorial in the Readme
- [ ] Add simple UI
- [x] Add logs
- [x] Add backup to BoltDB
- [x] Recover crashing instances
- [ ] Unit tests
