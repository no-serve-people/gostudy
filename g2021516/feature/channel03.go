package main

import (
	"fmt"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	channel := make(chan string) //注意: buffer为1
	go func() {
		channel <- "hello"
		fmt.Println("write \"hello\" done!")
		channel <- "world"
		fmt.Println("write \"World\" done!")
		fmt.Println("Write go sleep...")
		time.Sleep(3 * time.Second)
		channel <- "channel1"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Reader Wake up...")
	msg := <-channel
	fmt.Println("Reader: ", msg)
	msg = <-channel
	fmt.Println("Reader: ", msg)
	msg = <-channel //Writer在Sleep，这里在阻塞
	fmt.Println("Reader: ", msg)
}
