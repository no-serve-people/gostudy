package main

import (
	"fmt"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	// 设置一个timer，10钞后停掉ticker
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	ticker.Stop()
	fmt.Println("timer expired!")
}
