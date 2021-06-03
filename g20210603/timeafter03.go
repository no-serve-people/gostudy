package main

import (
	"fmt"
	"time"
)

// 难受，生产 Go timer.After 内存泄露之痛！
// https://juejin.cn/post/6969050288967647246
func main() {
	timer := time.NewTimer(3 * time.Minute)

	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())
		}
	}
}
