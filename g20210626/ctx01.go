package main

import (
	"context"
	"fmt"
)

// 说出一个避免Goroutine泄露的措施
// https://mp.weixin.qq.com/s/fl8WMsnfOxV0aD9Nic-F4w
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)

		if v == 5 {
			cancel()
			break
		}
	}
}
