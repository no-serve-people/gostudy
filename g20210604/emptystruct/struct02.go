package main

import "fmt"

// type T string
type T struct{}

func (c *T) Call() {
	fmt.Println("脑子进煎鱼了")
}

// https://mp.weixin.qq.com/s/8NJLe_11uvu22HfDjKqKjQ
// 详解 Go 空结构体的 3 种使用场景
func main() {
	var s T
	s.Call()
}
