package main

import (
	"os"
	"strings"
)
// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	os.Setenv("WEB", "https://coolshell.cn") //设置环境变量
	println(os.Getenv("WEB"))                //读出来

	for _, env := range os.Environ() { //穷举环境变量
		e := strings.Split(env, "=")
		println(e[0], "=", e[1])
	}
}
