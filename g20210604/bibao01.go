package main

import "fmt"

// https://mp.weixin.qq.com/s/OLgsdhXGEMltmjcpTW2ICw
// 一道 Go 闭包题，面试官说原来自己答错了：面别人也涨知识
func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}
