package main

import (
	"fmt"
)

type user02 struct {
	name string
	age  uint64
}

func main() {
	u := []user02{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user02, 0, len(u))
	for _, v := range u {
		o := v
		n = append(n, &o)
	}
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(v)
	}
}
