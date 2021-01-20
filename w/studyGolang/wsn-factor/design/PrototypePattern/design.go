package main

type Worker interface {
	Work(task *string)
}

type WorkerCreator interface {
	Create() Worker
}
