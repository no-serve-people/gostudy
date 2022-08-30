package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
//func IsClosed(ch <-chan T) bool {
//	select {
//	case <-ch:
//		return true
//	default:
//
//	}
//}

func main() {
	// c := make(chan T)
	// fmt.Println(IsClosed(c)) // false
	// close(c)
	// fmt.Println(IsClosed(c)) // true

	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumSenders = 1000
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}
	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Hour):
	}
}
