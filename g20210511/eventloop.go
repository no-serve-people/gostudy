package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func eventLoop(ctx context.Context) {
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
	// 建立一个可以手动取消的 Context
	ctx, cancel := context.WithCancel(context.Background())
	// 监控系统信号，这里只监控了 SIGINT（Ctrl+c），SIGTERM
	// 在 systemd 和 docker 中，都是先发 SIGTERM，过一段时间没退出再发 SIGKILL
	// 所以这里没捕获 SIGKILL
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	// 开始无限循环，收到信号就会退出
	eventLoop(ctx)
	fmt.Println("graceful shuwdown")
}
