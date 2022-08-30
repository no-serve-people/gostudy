package main

import (
	"fmt"
	"time"
)

// @see https://topgoer.cn/docs/goquestions/goquestions-1cjh2inun452c
// 解耦生产方和消费方
func main() {
	taskCh := make(chan int, 100)

	go worker3(taskCh)
	// 塞任务
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	// 等待一小时
	select {
	case <-time.After(time.Hour):
	}
}

func worker3(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
