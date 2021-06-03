package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		in := 1
		for {
			in++
			ch <- in
		}
	}()

	for {
		select {
		case _ = <-ch:
			// do something
			continue
		case <-time.After(3 * time.Minute):
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())
		}
	}
}
