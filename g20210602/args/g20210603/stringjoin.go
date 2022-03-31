package main

import (
	"fmt"
	"strings"
	"time"
)

// Go进阶---多使用strings.Join进行字符串拼接
// https://blog.csdn.net/weixin_37720172/article/details/104337362
func main() {
	num := 50000
	list := make([]string, 100)

	for i := 0; i < num; i++ {
		list = append(list, "helloworld")
	}

	var test1 string
	var seq string

	start1 := time.Now()
	for i := 0; i < num; i++ {
		test1 += seq + list[i]
		seq = " "
	}
	end1 := time.Now().Sub(start1)
	start2 := time.Now()
	test2 := strings.Join(list, " ")
	end2 := time.Now().Sub(start2)

	fmt.Println(test1, test2) // 输出下两个字符串
	fmt.Println(end1)
	fmt.Println(end2)
}
