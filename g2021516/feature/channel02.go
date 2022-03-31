package main

import "fmt"

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	channel1 := make(chan string, 2)
	go func() {
		channel1 <- "hello"
		channel1 <- "world"
	}()

	msg1 := <-channel1
	msg2 := <-channel1

	fmt.Println(msg1, msg2)
}
