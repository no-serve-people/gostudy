package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
var total_tickets int32 = 10
var mutex = &sync.Mutex{} // 可简写成：var mutex sync.Mutex

func sell_tickets(i int) {
	//for {
	//	if total_tickets > 0 { // 如果有票就卖
	//		mutex.Lock()
	//		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	//		total_tickets-- //卖一张票
	//		fmt.Println("id:", i, " ticket:", total_tickets)
	//	} else {
	//		break
	//	}
	//}
	for total_tickets > 0 {
		mutex.Lock()
		if total_tickets > 0 { // 如果有票就卖
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets-- // 卖一张票
			fmt.Println("id:", i, " ticket:", total_tickets)
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(4)        // 我的电脑是4核处理器，所以我设置了4
	rand.Seed(time.Now().Unix()) // 生成随机种子

	for i := 0; i < 5; i++ {
		go sell_tickets(i)
	}
	// 等待线程执行完
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done") // 退出时打印还有多少票
}
