package main

// 输出子协程结果方式01
// @fixme 方式2 请看 memory.go
import (
	"fmt"
	"time"
)

func add02(a, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	go add02(1, 2)
	time.Sleep(1e9) // 等待 1s
}

// 要显示出子协程的打印结果，第一种方式是在主协程中等待足够长的时间再退出，以保证子协程中的所有代码执行完毕：
// @fixme 存在的问题 对于这种加法操作 等待1s肯定够了 但是如果对于 发邮件 数据库连接之类的操作 很难预估时间的  这种方式将不奏效
// @fixme 方式2 请看 memory.go
