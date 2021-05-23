package main

import "fmt"

// https://studygolang.com/articles/28006#reply1
// 指针的详细讲解
func main() {
	var x int = 99
	var p *int = &x
	fmt.Println(p)

	x = 100

	fmt.Println("x: ", x)
	fmt.Println("*p", *p)

	*p = 99
	fmt.Println("x ", x)
	fmt.Println("*p: ", *p)
}
