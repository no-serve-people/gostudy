package main

import "fmt"

type user04 struct {
	name string
	age  uint64
}

func main() {
	u := []user04{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	for k, v := range u {
		if v.age != 18 {
			u[k].age = 18
		}
	}

	fmt.Println(u)
}
