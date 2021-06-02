package main

import (
	"fmt"
	"runtime"
)

func example1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(Stack()))
		}
	}()
	panic("unknown")
}

func example2() {
	defer recover()
	panic("unknown")
}
func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
