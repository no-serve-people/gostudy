package main

import (
	"context"
	"time"
)

// https://github.com/cch123/golang-notes/blob/master/context.md
// context

func main() {
	ctx, _ := context.WithCancel(context.TODO())

	for i := 0; i < 100; i++ {
		go func() {
			select {
			case <-ctx.Done():
			}
		}()
	}
	time.Sleep(time.Hour)
}
