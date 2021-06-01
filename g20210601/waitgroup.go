package main

import (
	"fmt"
	"sync"
)

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484551&idx=1&sn=aff008de8b6c089d7daf9c963ce8b60a&scene=21#wechat_redirect
// 源码剖析sync.WaitGroup(文末思考题你能解释一下吗?)

type httpPkg struct {
}

var http httpPkg

func (httpPkg) Get(url string) {

}

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func exampleImplWaitGroup() {
	done := make(chan struct{}) // 收10份保护费

	count := 10 // 10个马仔
	for i := 0; i < count; i++ {
		go func(i int) {
			defer func() {
				done <- struct{}{}
			}()
			fmt.Printf("马仔%d号收保护费\n", i)
		}(i)
	}

	for i := 0; i < count; i++ {
		<-done
		fmt.Printf("马仔%d号已经收完保护费\n", i)
	}
	fmt.Println("所有马仔已经干完活了，开始酒吧消费～")
}
