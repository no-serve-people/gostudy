package main

import "fmt"

func main() {
	var args int64 = 1
	modifiedNumber(args) // args就是实际参数
	fmt.Printf("实际参数的地址 %p\n", &args)
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumber(args int64) { // 这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &args)
	args = 10
}
