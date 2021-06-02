package main

import "fmt"

type user03 struct {
	name string
	age  uint64
}

func main() {
	u := []user03{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	for _, v := range u {
		if v.age != 18 {
			v.age = 20
		}
	}
	fmt.Println(u)
}
