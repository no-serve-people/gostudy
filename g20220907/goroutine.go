package g20220907

import (
	"context"
	"fmt"
	"time"
)

// https://topgoer.cn/docs/goquestions/goquestions-1cjh3lkbt937u
// context 作用
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}
