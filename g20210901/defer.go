package main

import "fmt"

func main() {
	fmt.Println(do())
}

func do() string {
	defer fmt.Println("do defer")
	fmt.Println("do something")
	// panic("do panic")

	return "do return"
}
