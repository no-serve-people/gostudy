package main

import "fmt"

// https://studygolang.com/articles/28006#reply1
// 指针的详细讲解
func swap(x, y int) {
	x, y = y, x
	fmt.Println("swap  x: ", x, "y: ", y)
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 20
	// swap(x, y)
	swap2(&x, &y)
	fmt.Println("main  x: ", x, "y: ", y)
}
