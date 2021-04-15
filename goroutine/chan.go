package main

import (
	"fmt"
	"time"
)

// 协程通信-消息传递 通过通道
func add5(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1 // 表示把元素 1 发送到通道 ch
}

func main() {
	start := time.Now()
	// 定义一个包含 10 个通道类型的切片 chs
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int) // 定义通道chan 是一种数据类型 一个通道只能传递一种类型的值 这里指定 int
		go add5(1, i, chs[i])
	}
	// 等所有子协程执行完毕后 才会执行
	for _, ch := range chs {
		<-ch // 接收通道 指定的变量忽略
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

// 往通道写入数据和从通道接收数据都是原子操作，是同步阻塞的
// 通道的发送和接收操作是互斥的 同一时间同一个进程内的所有协程对某个通道只能执行发送或接收操作 这样就保证了并发的安全性，数据不可能被污染。

// @fixme 程序耗时和共享内存相当，但是代码要简洁的多，优雅的多。
// @fixme 不过 go 语言还是存在不少需要 共享内存+锁 方式实现并发编程的情况 这个后续再讲
