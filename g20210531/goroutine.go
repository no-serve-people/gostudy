package main

import (
	"fmt"
	"sync"
	"time"
)

// https://juejin.cn/post/6967737755476426765
// Golang 并发编程（上）｜周末学习
func main() {
	// ============== sync ==============
	startTime1 := time.Now()
	testGoroutine()

	endTime1 := time.Now()
	// 可以看到串行需要 3s 的下载操作，并发后，只需要 1s
	fmt.Printf("======> Done! use time: %f", endTime1.Sub(startTime1).Seconds())
	fmt.Println()

	// ============== channel  ==============

	startTime2 := time.Now()

	testChannel()

	endTime2 := time.Now()
	fmt.Printf("======> Done! use time: %f", endTime2.Sub(startTime2).Seconds())
	fmt.Println()
}

// 使用 sync.WaitGroup 多个并发协程之间不能通信, 等待所有并发协程执行结束
var wg sync.WaitGroup

func testGoroutine() {
	for i := 0; i < 3; i++ {
		wg.Add(1)                                    // wg.Add() 为 wg 添加一个计数; wg.Done() 减去一个计数。
		go downloadV1("a.com/time_" + string(i+'0')) // 启动新的协程并发执行 download 函数。
	}
	wg.Wait() // wg.Wait()：等待所有的协程执行结束。
}

// 使用 channel 信道，可以在协程之间传递消息。等待并发协程返回消息。
var ch = make(chan string, 10) // 创建大小 10 的缓存通道
func testChannel() {
	for i := 0; i < 3; i++ {
		go downloadV2("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
}

// ================= 工具 V  =================
func downloadV1(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时
	wg.Done()
}

func downloadV2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}
