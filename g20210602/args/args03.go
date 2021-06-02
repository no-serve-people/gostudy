package main

import "fmt"

func main() {
	var args int64 = 1

	addr := &args
	fmt.Printf("原始指针的内存地址是 %p\n", addr)
	modifiedNumber02(addr) // args就是实际参数
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumber02(addr *int64) {
	fmt.Printf("形参地址 %p \n", &addr)
	*addr = 10
}
