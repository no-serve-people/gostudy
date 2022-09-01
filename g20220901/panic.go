package g20220901

import "fmt"

func testPanic() {
	// 错误的示例
	recover()
	panic("not good") // 发生 panic，主程序退出
	recover()         // 不会被执行
	println("ok")
}

func testPanic2() {
	// 正确的示例
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	panic("not google")
}
