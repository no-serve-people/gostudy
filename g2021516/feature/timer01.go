package main

import (
	"fmt"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	fmt.Println("timer expired!")
}
