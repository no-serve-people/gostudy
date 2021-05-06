package limit

// https://my.oschina.net/qiangmzsx/blog/4277685 不得不了解系列之限流
import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Counter struct {
	Count       uint64 // 初始计数器
	Limit       uint64 // 单位时间窗口最大请求频次
	Interval    int64  // 单位ms
	RefreshTime int64  // 时间窗口
}

func NewCounter(count, limit uint64, interval, rt int64) *Counter {
	return &Counter{
		Count:       count,
		Limit:       limit,
		Interval:    interval,
		RefreshTime: rt,
	}
}

func (c *Counter) RateLimit() bool {
	now := time.Now().UnixNano() / 1e6
	if now < (c.RefreshTime + c.Interval) {
		atomic.AddUint64(&c.Count, 1)
		return c.Count <= c.Limit
	} else {
		c.RefreshTime = now
		atomic.AddUint64(&c.Count, -c.Count)
		return true
	}
}

func Test_Counter(t *testing.T) {
	counter := NewCounter(0, 5, 100, time.Now().Unix())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for k := 0; k <= 10; k++ {
				fmt.Println(counter.RateLimit())
				if k%3 == 0 {
					time.Sleep(102 * time.Millisecond)
				}
			}
		}(i)
		time.Sleep(10 * time.Second)
	}
}
