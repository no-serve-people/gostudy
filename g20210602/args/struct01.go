package main

import "fmt"

type Person struct {
	Name string
	Age  int64
}

func main() {
	per := Person{
		Name: "asong",
		Age:  8,
	}
	fmt.Printf("原始struct地址是：%p\n", &per)
	modifiedAge02(per)
	fmt.Println(per)
}
func modifiedAge02(per Person) {
	fmt.Printf("函数里接收到struct的内存地址是：%p\n", &per)
	per.Age = 10
}
