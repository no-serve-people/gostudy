package main

import "fmt"

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484131&idx=1&sn=5b1608613f6f1568acc85214e0ff1cbd&scene=21#wechat_redirect
// 女朋友问我：小松子，你知道Go语言参数传递是传值还是传引用吗？
func main() {
	var args int64 = 1
	println(args) // args就是实际参数
}

func printNumber(args ...int64) { // 这里定义的args就是形式参数
	for _, arg := range args {
		fmt.Println(arg)
	}
}
