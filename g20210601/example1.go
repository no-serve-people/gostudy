package main

import (
	"fmt"
	"runtime"
	"time"
)

// 源码剖析panic与recover，看不懂你打我好了！
// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484460&idx=1&sn=0838818da3a456f409f70dd321ffa2e6&scene=21#wechat_redirect
func example1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(Stack()))
		}
	}()
	panic("unknown")
}

func example2() {
	defer recover()
	panic("unknown")
}

func example3() {
	defer fmt.Println("this is a example3 for defer use panic")
	defer func() {
		defer func() {
			panic("panic defer 2")
		}()
		panic("panic defer 1")
	}()
	panic("panic example3")
}

func example4() {
	fmt.Println("goroutine example4")
	defer func() {
		fmt.Println("test defer")
	}()
	panic("unknown")
}

func example5() {
	defer fmt.Println("goroutine example5")
	time.Sleep(5 * time.Second)
}

func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

func main() {
	example1()
	example2()
}
