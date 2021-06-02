package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484475&idx=1&sn=f470d42fee8b9b65fad7b43e57518cfc&scene=21#wechat_redirect
// 详解并发编程基础之原子操作(atomic包)
func main() {
	var share uint64 = 1
	wg := sync.WaitGroup{}

	wg.Add(3)
	// 协程1，期望值是1,欲更新的值是2
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapUint64(&share, 1, 2)
		fmt.Println("goroutine 1", swapped)
	}()
	// 协程2，期望值是1，欲更新的值是2
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		swapped := atomic.CompareAndSwapUint64(&share, 1, 2)
		fmt.Println("goroutine 2", swapped)
	}()
	// 协程3，期望值是2，欲更新的值是1
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		swapped := atomic.CompareAndSwapUint64(&share, 2, 1)
		fmt.Println("goroutine 3", swapped)
	}()

	wg.Wait()

	fmt.Println("main exit")

}
