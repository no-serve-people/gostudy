package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// 通过 context 包提供的函数实现多协程之间的协作
// https://laravelacademy.org/post/20005.html
func AddNum(a *int32, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(a)
		newNum := curNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number当前值: %d [%d-%d]\n", *a, b, i)
			break
		}
	}
}

func main() {
	total := 0
	var num int32
	fmt.Printf("number初始值: %d\n", num)
	fmt.Println("启动子协程...")

	ctx, CancelFunc := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go AddNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				CancelFunc()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("所有子协程执行完毕.")
}
