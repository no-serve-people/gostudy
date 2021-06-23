package main

import (
	"fmt"
	"sync"
	"time"
)

var group int

var cond *sync.Cond

func test1() {
	fmt.Println("任务1")

	for i := 0; i < 5; i++ {
		group = i
		cond.Broadcast()
		time.Sleep(time.Millisecond * 1)
	}

}

func test2() {
	for group != 2 {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
	}
	fmt.Println("任务2")
}

func test3() {
	for group != 3 {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
	}
	println("任务3")
}

func do(fs ...func()) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	for _, f := range fs {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}

	return &wg
}
func main() {
	cond = sync.NewCond(&sync.Mutex{})
	wg := do(test1, test2, test3)

	wg.Wait()

}
