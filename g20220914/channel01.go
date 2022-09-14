package g20220914

import "fmt"

// 测试无缓冲和有缓冲 读取情况
func testChannel() {
	// var a = make(chan int, 5)
	a := make(chan int)
	// a <- 2
	close(a)

	value, ok := <-a
	if ok {
		fmt.Println(value, "呵呵1")
	}

	value, ok = <-a
	if !ok {
		fmt.Println(value, "channel closed, data invalid.")
	}
}
