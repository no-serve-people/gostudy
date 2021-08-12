package main

import (
	"fmt"
	"gostudy/g20210805/workpool"
	"time"
)

// Go语言的并发与WorkerPool - 第二部分
// https://mp.weixin.qq.com/s/u3h8mfNja_SAnFP4vawCGg
func main() {
	var allTask []*workpool.Task
	for i := 0; i < 100; i++ {
		task := workpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workpool.NewPool(allTask, 5)
	pool.Run()
}
