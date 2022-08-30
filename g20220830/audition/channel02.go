package main

import (
	"fmt"
	"time"
)

// @see https://topgoer.cn/docs/goquestions/goquestions-1cjh2jsk14499
func worker() {
	ticker := time.Tick(1 * time.Second)

	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}
