package main

import (
	"fmt"
	"sync"
)

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484505&idx=1&sn=73ab61b21574d6ed1c11b6a5516888b0&scene=21#wechat_redirect
// 详解并发编程之sync.Once的实现(附上三道面试题)
func panicDo() {
	once := &sync.Once{}
	defer func() {
		if err := recover(); err != nil {
			once.Do(func() {
				fmt.Println("run in recover")
			})
		}
	}()

	once.Do(func() {
		panic("panic i=0")
	})
}

func nestedDo() {
	once := &sync.Once{}
	once.Do(func() {
		once.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}
func nestedDo02() {
	once1 := &sync.Once{}
	once2 := &sync.Once{}
	once1.Do(func() {
		once2.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}

func main() {
	panicDo()
	nestedDo()
	nestedDo02()
}
