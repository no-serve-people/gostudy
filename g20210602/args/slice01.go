package main

import "fmt"

func main() {
	args := []int64{1, 2, 3}
	fmt.Printf("切片args的地址： %p \n", args)
	fmt.Printf("切片args第一个元素的地址： %p \n", &args[0])
	fmt.Printf("直接对切片args取地址%v \n", &args)
	fmt.Printf("切片args的地址： %p\n", args)
	modifiedNumber03(args)
	fmt.Println(args)
}

func modifiedNumber03(args []int64) {
	fmt.Printf("形参切片的地址 %p \n", args)
	fmt.Printf("形参切片args第一个元素的地址： %p \n", &args[0])
	fmt.Printf("直接对形参切片args取地址%v \n", &args)
	args[0] = 10
}
