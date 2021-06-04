package main

import (
	"fmt"
	"time"
)

// https://mp.weixin.qq.com/s/8NJLe_11uvu22HfDjKqKjQ
// 详解 Go 空结构体的 3 种使用场景
func main() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	fmt.Println("脑子好像进...")
	<-ch
	fmt.Println("煎鱼了！")
}
