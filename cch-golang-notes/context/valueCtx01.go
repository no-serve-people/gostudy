package main

import (
	"context"
	"fmt"
)

// https://github.com/cch123/golang-notes/blob/master/context.md
// context
type orderId int

func main() {
	x := context.TODO()
	x = context.WithValue(x, orderId(1), "1234")
	x = context.WithValue(x, orderId(2), "2345")
	x = context.WithValue(x, orderId(3), "3456")
	fmt.Println(x.Value(orderId(2)))
}
