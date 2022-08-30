package main

import (
	"fmt"
	"time"
)

// @see https://topgoer.cn/docs/goquestions/goquestions-1cjh2ok9b2vn1
func goroutineC(a <-chan int) {
	val := <-a
	fmt.Println("G1 received data: ", val)
	return
}

func goroutineD(b <-chan int) {
	val := <-b
	fmt.Println("G2 received data: ", val)
	return
}

func main() {
	ch := make(chan int)
	go goroutineC(ch)
	go goroutineD(ch)
	ch <- 3
	time.Sleep(time.Second)
}
