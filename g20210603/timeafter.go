package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- "脑子进煎鱼了"
	}()

	select {
	case _ = <-ch:
	case <-time.After(time.Second * 1):
		fmt.Println("煎鱼出去了，超时了！！！")
	}
}
