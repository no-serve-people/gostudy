package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	var cnt uint32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&cnt, 1)
			}
		}()
	}

	time.Sleep(time.Second)             // 等一秒钟等goroutine完成
	cntFinal := atomic.LoadUint32(&cnt) // 取数据
	fmt.Println("cnt:", cntFinal)
}
