package main

import (
	"sync"
	"time"
)

// https://mp.weixin.qq.com/s/WvlVQuyMsPQzWJqBW6qQ_g
// Go：无缓冲和有缓冲通道
func main() {
	c := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- "foo"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println("Message" + <-c)
	}()

	wg.Wait()
}
