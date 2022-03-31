package main

import (
	"fmt"
	"net"
	"time"
)

const RECV_BUF_LEN_CLIENT = 1024

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN_CLIENT)

	for i := 0; i < 5; i++ {
		// 准备要发送的字符串
		msg := fmt.Sprintf("hello world %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		fmt.Println(msg)
		// 从服务器端收字符串
		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		fmt.Println(string(buf[0:n]))
		// 等一秒钟
		time.Sleep(time.Second)
	}
}
