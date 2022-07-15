package queue

import (
	jobqueue "github.com/dirkaholic/kyoo"
	"runtime"
)

var Queue *jobqueue.JobQueue

func StartQueue(){
	Queue = jobqueue.NewJobQueue(runtime.NumCPU() * 2)
	Queue.Start()
}