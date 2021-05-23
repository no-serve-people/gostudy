package main

import (
	"fmt"
	"strconv"
	"sync"
)

// https://studygolang.com/topics/13680#reply3
// 请教一个goroutine报错问题？

var wg1 = sync.WaitGroup{}

var chexit = make(chan bool)

func gen(name string) chan int {
	ch := make(chan int)
	wg1.Add(1)
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case ch <- i:
			case <-chexit:
				close(ch)
				fmt.Println("go routine exit:", name)
				wg1.Done()
				return
			}
		}
	}()
	return ch
}

func primeFilter(chin chan int, prime int, name string) chan int {
	chout := make(chan int)

	wg1.Add(1)

	go func() {
		for {
			select {
			case v := <-chin:
				if v%prime != 0 {
					chout <- v
				}

			case <-chexit:
				close(chout)
				fmt.Println("go routine exit:", name)
				wg1.Done()
				return
			}
		}
	}()
	return chout
}

func Prime() {
	ch := gen("generate numbers")

	for i := 0; i < 100; i++ {
		v := <-ch
		// log.Printf("%d,%d\n", i+1, v)
		ch = primeFilter(ch, v, "filter "+strconv.Itoa(i))
	}
	close(chexit)
	wg1.Wait()
}
