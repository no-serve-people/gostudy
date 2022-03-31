package main

import "fmt"

func main() {
	array := []int{7, 8, 9}
	fmt.Printf("main ap brfore: len: %d cap:%d data:%+v\n", len(array), cap(array), array)
	ap(array)
	fmt.Printf("main ap after: len: %d cap:%d data:%+v\n", len(array), cap(array), array)
}

func ap(array []int) {
	fmt.Printf("ap brfore:  len: %d cap:%d data:%+v\n", len(array), cap(array), array)
	array[0] = 1
	array = append(array, 10)
	fmt.Printf("ap after:   len: %d cap:%d data:%+v\n", len(array), cap(array), array)
}
