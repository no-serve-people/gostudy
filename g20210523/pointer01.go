package main

import "fmt"

// https://studygolang.com/articles/28006#reply1
// 指针的详细讲解
func main() {
	var x int = 10
	var p *int = &x

	fmt.Println(p)
}
