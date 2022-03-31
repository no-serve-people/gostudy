package main

import (
	"context"
	"fmt"
)

// https://github.com/cch123/golang-notes/blob/master/context.md
// context
type orderId2 int

func main() {
	x := context.TODO()
	x = context.WithValue(x, orderId2(1), "1234")
	x = context.WithValue(x, orderId2(2), "2345")

	y := context.WithValue(x, orderId2(3), "4567")
	x = context.WithValue(x, orderId2(3), "3456")

	fmt.Println(x.Value(orderId2(3)))
	fmt.Println(y.Value(orderId2(3)))
}
