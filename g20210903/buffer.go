package main

import (
	"sync"
	"time"
)
// https://mp.weixin.qq.com/s/WvlVQuyMsPQzWJqBW6qQ_g
// Go：无缓冲和有缓冲通道
func main() {
	c := make(chan string, 2)

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		c <- `foo`
		c <- `bar`
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		println(`Message: ` + <-c)
		println(`Message: ` + <-c)
	}()
}
