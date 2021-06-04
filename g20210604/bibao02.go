package main

import "fmt"

// https://mp.weixin.qq.com/s/OLgsdhXGEMltmjcpTW2ICw
// 一道 Go 闭包题，面试官说原来自己答错了：面别人也涨知识
func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
