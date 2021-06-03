package main

import (
	"fmt"
)

// https://www.v2ex.com/t/781192#reply6
// go 语言的一个死锁问题
func main() {
	// 创建 3 个 channel,A,B 和 Exit
	A := make(chan bool)
	B := make(chan bool)
	Exit := make(chan bool)

	go func() {
		// 如果 A 通道是 true,我就执行
		for i := 0; i < 10; i++ {
			if ok := <-A; ok {
				fmt.Println("A 输出", i)
				B <- true
			}
		}
	}()

	go func() {
		defer func() {
			Exit <- true // 这个协程的活干完之后，向主 goroutine 发送信号
		}()
		// 如果 B 通道是 true,我就执行
		for i := 2; i < 10; i += 2 {
			if ok := <-B; ok {
				fmt.Println("B 输出", i)
				if i != 10 { // 如果 i 等于 10 了，就不要再向 A 通道写数据了，否则将导致 A 通道死锁，至于为什么，坦白说我很疑惑
					A <- true
				}
			}
		}
	}()
	A <- true // 启动条件
	<-Exit    // 结束条件
}
