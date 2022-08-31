package main

import "fmt"

func main() {
	ageMap := make(map[string]int)
	ageMap["qcrao"] = 18
	// 不带comma的用法
	age1 := ageMap["stefno"]
	fmt.Println(age1)
	// 带comma的用法
	age2, ok := ageMap["stefno"]
	fmt.Println(age2, ok)
}
