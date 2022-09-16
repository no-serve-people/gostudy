package g20220915

import (
	"fmt"
	"sync"
	"time"
)

// 多线程卖票，100张票5个线程，保证不要超卖，尽可能使所有线程都在工作不要空闲

// 最好使用cahnnel
// 使用mutex
func testMutex() {
	n := 100

	var wg sync.WaitGroup

	var mutex sync.RWMutex
	wg.Add(5)

	for i := 0; i < 5; i++ {
		// fmt.Println(i,"现在是")
		go func(i int) {
			defer wg.Done()
			for {
				mutex.Lock()
				fmt.Printf("协程%d买到一张票\n", i)
				n--
				if n == 0 {
					fmt.Println("票卖完了")
					return
				}
				mutex.Unlock()
			}
		}(i)
	}

	wg.Wait()

	// fmt.Println(n)
}

// 使用channel
func testChannel() {
	ch := make(chan int, 100)
	for j := 100; j > 0; j-- {
		ch <- j
	}

	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				t := <-ch
				if t > 0 {
					fmt.Printf("买到票了，协程id%d 买到票%d \n", i, t)
				} else {
					fmt.Printf("票卖完了,协程id%d \n", i)
					break
				}
			}
		}(i)
	}

	time.Sleep(5 * time.Second)
}
