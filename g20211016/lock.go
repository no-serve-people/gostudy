package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 如何用硬件同步原语（CAS）替代锁？
// http://learn.lianglianglee.com/%E4%B8%93%E6%A0%8F/%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97%E9%AB%98%E6%89%8B%E8%AF%BE/18%20%20%E5%A6%82%E4%BD%95%E7%94%A8%E7%A1%AC%E4%BB%B6%E5%90%8C%E6%AD%A5%E5%8E%9F%E8%AF%AD%EF%BC%88CAS%EF%BC%89%E6%9B%BF%E4%BB%A3%E9%94%81%EF%BC%9F.md
func main() {
	// 账户初始值为 0 元
	var balance int32
	balance = int32(0)
	done := make(chan bool)
	// 执行 10000 次转账，每次转入 1 元
	count := 10000
	var lock sync.Mutex

	for i := 0; i < count; i++ {
		// 这里模拟异步并发转账
		go transfer(&balance, 1, done, &lock)
		// go transferCas(&balance, 1, done)
	}
	// 等待所有转账都完成
	for i := 0; i < count; i++ {
		<-done
	}
	// 打印账户余额
	fmt.Println("balace=%d\n", balance)
}

// 转账服务
func transfer(balance *int32, amount int, done chan bool, lock *sync.Mutex) {
	lock.Lock()
	*balance += int32(amount)
	lock.Unlock()
	done <- true
}

func transferCas(balance *int32, amount int, done chan bool) {
	for {
		old := atomic.LoadInt32(balance)
		new := old + int32(amount)

		if atomic.CompareAndSwapInt32(balance, old, new) {
			break
		}
	}
	done <- true
}
