package main

import (
	"context"
	"fmt"
	"os/signal"
	"time"
)

func eventLoop02(ctx context.Context) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			// 收到信号退出无限循环
			break LOOP
		default:
			// 用一个 sleep 模拟业务逻辑
			time.Sleep(time.Second * 10)
		}
	}
}
func main() {
	// 监控系统信号和创建 Context 现在一步搞定
	ctx, stop := signal.NotifyContext(context.Background())
	// 在收到信号的时候，会自动触发 ctx 的 Done ，这个 stop 是不再捕获注册的信号的意思，算是一种释放资源。

	defer stop()
	// 开始无限循环，收到信号就会退出
	eventLoop02(ctx)
	fmt.Println("graceful shuwdown")
}
