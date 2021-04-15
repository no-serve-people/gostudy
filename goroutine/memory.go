package main

// 协程通信-共享内存（实现并发：共享内存+锁）
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0 // 初始化 该共享内存被所有协程共享

func add04(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++ // 写操作的并发
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	//runtime.GOMAXPROCS(1)
	start := time.Now()
	lock := &sync.Mutex{} // 互斥锁
	for i := 0; i < 10; i++ {
		go add04(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter // 读操作的并发
		lock.Unlock()
		runtime.Gosched() // 用于让出CPU时间片
		if c >= 10 {
			break // 退出循环
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

// 主线程和子线程 都加锁 是为了防止 读写并发问题  保证对 counter 进行读取和更新操作时，同时只有一个协程在操作它（既保证了操作的原子性）

// @fixme 问题 繁琐  需要加锁 解锁 容易出错
