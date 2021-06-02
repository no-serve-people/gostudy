package main

import "fmt"

func main() {
	persons := make(map[string]int)
	persons["asong"] = 8

	addr := &persons
	fmt.Printf("原始map的内存地址是：%p\n", addr)
	modifiedAge(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func modifiedAge(person map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &person)
	person["asong"] = 9
}
