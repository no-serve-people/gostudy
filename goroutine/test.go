package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Second)
		}()

	}
	wg.Wait() // 等待在此，等所有go func里都执行了Done()才会退出
}
