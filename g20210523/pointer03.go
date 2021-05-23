package main

import "fmt"

// https://studygolang.com/articles/28006#reply1
// 指针的详细讲解
func main() {
	p := new(int)

	*p = 1000

	fmt.Println(p)
	fmt.Println(*p)
}
