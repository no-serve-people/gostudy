package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484505&idx=1&sn=73ab61b21574d6ed1c11b6a5516888b0&scene=21#wechat_redirect
// 详解并发编程之sync.Once的实现(附上三道面试题)
type MyOnce struct {
	flag uint32
	lock sync.Mutex
}

func (m *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&m.flag) == 0 {
		m.lock.Lock()
		defer m.lock.Unlock()

		if atomic.CompareAndSwapUint32(&m.flag, 0, 1) {
			f()
		}
	}
}

func testDo() {
	mOnce := MyOnce{}

	for i := 0; i < 10; i++ {
		go func() {
			mOnce.Do(func() {
				fmt.Println("test my once only run once")
			})
		}()
	}
}

func main() {
	testDo()
	time.Sleep(10 * time.Second)
}
