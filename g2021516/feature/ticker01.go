package main

import (
	"fmt"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	ticker := time.NewTicker(time.Second)

	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
