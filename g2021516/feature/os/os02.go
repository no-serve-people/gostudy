package main

import (
	"fmt"
	"os/exec"
)

// https://coolshell.cn/articles/8489.html
// GO 语言简介（下）— 特性
func main() {
	cmd := exec.Command("ping", "127.0.0.1")

	out, err := cmd.Output()
	if err != nil {
		println("Command Error!", err.Error())
		return
	}

	fmt.Println(string(out))
}
