package g20220922

import "fmt"

// golang学习--关于defer的执行顺序

func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

func deferCall2(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func deferCall3() {
	a := 1
	b := 2
	defer deferCall2("1", a, deferCall2("10", a, b))
	a = 0
	defer deferCall2("2", a, deferCall2("20", a, b))
	b = 1
}
