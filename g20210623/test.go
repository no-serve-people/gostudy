//package main
//
//import "fmt"

//
//func f1(in chan int) {
//	fmt.Println(<-in)
//	//fmt.Println(111)
//}
//
//func main() {
//	out := make(chan int, 1)
//	out <- 2
//	go f1(out)
//
//	time.Sleep(time.Second)
//}
//package main
//
//import (
//"fmt"
//)
//
//func f1(in chan int) {
//	fmt.Println(<-in)
//}
//
//func main() {
//	out := make(chan int)
//	go f1(out)
//	out <- 2
//}
package main

import "time"

var (
	c = make(chan int)
	a string
)

func f() {
	a = "hello, world"
	<-c
}

// happen before
func main() {
	c <- 0
	go f()
	print(a)
	time.Sleep(2 * time.Second)
}
