package main

import (
	"fmt"
	"unsafe"
)
// https://mp.weixin.qq.com/s/8NJLe_11uvu22HfDjKqKjQ
// 详解 Go 空结构体的 3 种使用场景
func main() {
	var a int
	var b string
	var c bool
	var d [3]int32
	var e []string
	var f map[string]bool
	var s struct{}
	fmt.Println(
		unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c),
		unsafe.Sizeof(d),
		unsafe.Sizeof(e),
		unsafe.Sizeof(f),
		unsafe.Sizeof(s),
	)
}
