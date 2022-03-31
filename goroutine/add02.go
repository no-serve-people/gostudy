package main

// 验证协程是并发执行的
import (
	"fmt"
	"time"
)

func add03(a, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(3e9) // 休眠3秒 遇到阻塞 进行协程切换
}

func main() {
	for i := 0; i < 10; i++ {
		go add03(1, i)
	}
	time.Sleep(1e9) // 等待1s
}

// 结果不是按照严格的意义的顺序执行的 而是乱序的
