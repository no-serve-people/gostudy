package main
// https://mp.weixin.qq.com/s/szSxatDakPQMUA8Vm9u3qQ
// 源码剖析sync.cond(条件变量的实现机制）
import (
	"fmt"
	"sync"
	"time"
)

var (
	done  = false
	topic = "Golang梦工厂"
)

func Consumer(topic string, cond *sync.Cond) {
	cond.L.Lock()
	for !done {
		cond.Wait()
	}
	fmt.Println("topic is ", topic, " starts Consumer")
	cond.L.Unlock()
}

func Push(topic string, cond *sync.Cond) {
	fmt.Println(topic, "starts Push")
	cond.L.Lock()
	done = true

	cond.L.Unlock()
	fmt.Println("topic is ", topic, " wakes all")
	cond.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go Consumer(topic, cond)
	go Consumer(topic, cond)
	go Consumer(topic, cond)

	Push(topic, cond)
	time.Sleep(5 * time.Second)
}
