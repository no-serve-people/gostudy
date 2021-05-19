package main

import (
	"context"
	"fmt"
	"time"
)

// https://github.com/cch123/golang-notes/blob/master/context.md
// context

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

	defer cancel()

	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():

		fmt.Println(ctx.Err())
	}
}
