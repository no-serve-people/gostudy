package main

import (
	"fmt"
	"strings"
)

func checkZeroOnes(s string) bool {
	zero, one := 0, 0
	for _, v := range strings.Split(s, "1") {
		zero = Max(len(v), zero)
	}
	for _, v := range strings.Split(s, "0") {
		one = Max(len(v), one)
	}
	return one > zero
}

func Max(res int, nums ...int) int {
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	s := "1101"
	fmt.Println(checkZeroOnes(s))
}
