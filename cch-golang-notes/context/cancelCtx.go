package main

import (
	"context"
	"fmt"
	"time"
)

// https://github.com/cch123/golang-notes/blob/master/context.md
// context

func main() {
	jobChan := make(chan struct{})
	ctx, cancelFn := context.WithCancel(context.TODO())

	worker := func() {
	jobLoop:
		for {
			select {
			case <-jobChan:
				fmt.Println("do my job")
			case <-ctx.Done():
				fmt.Println("parent call me to quit")
				break jobLoop
			}
		}
	}

	// start worker
	go worker()
	// stop all worker
	cancelFn()
	time.Sleep(time.Second)
}
