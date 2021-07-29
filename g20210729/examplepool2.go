package main

import (
	"fmt"
	"gostudy/g20210729/wokerpool2"
)

func main() {
	pool := &wokerpool2.Pool{
		Name:      "test",
		Size:      5,
		QueueSize: 20,
	}
	pool.Initialize()
	pool.Start()

	defer pool.Stop()
	for i := 0; i < 100; i++ {
		job := &PrintJob{
			Index: i,
		}
		pool.Queue <- job
	}
}

// PrintJob ...
type PrintJob struct {
	Index int
}

func (pj *PrintJob) Start(worker *wokerpool2.Worker) error {
	fmt.Printf("job %s - %d\n", worker.Name, pj.Index)
	return nil
}
