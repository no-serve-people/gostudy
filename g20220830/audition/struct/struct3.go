package _struct

import "fmt"

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func testChannel() {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}
