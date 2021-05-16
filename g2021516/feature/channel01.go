package main

import "fmt"
// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	// 创建一个string类型的channel
	channel := make(chan string)
	// 创建一个goroutine向channel里发一个字符串
	go func() { channel <- "hello" }()
	msg := <-channel
	fmt.Println(msg)
}
