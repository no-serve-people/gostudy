package main

import (
	"fmt"
	"sync"
)

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484551&idx=1&sn=aff008de8b6c089d7daf9c963ce8b60a&scene=21#wechat_redirect
// 源码剖析sync.WaitGroup(文末思考题你能解释一下吗?)
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
