package main

import (
	"fmt"
	"time"
)

// 女朋友问我：小松子，你知道Go语言参数传递是传值还是传引用吗？
// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484131&idx=1&sn=5b1608613f6f1568acc85214e0ff1cbd&scene=21#wechat_redirect
func main() {
	p := make(chan bool)
	fmt.Printf("原始chan的内存地址是：%p\n", &p)

	go func(p chan bool) {
		fmt.Printf("函数里接收到chan的内存地址是：%p\n", &p)
		// 模拟耗时
		time.Sleep(2 * time.Second)
		p <- true
	}(p)

	select {
	case l := <-p:
		fmt.Println(l)
	}
}
