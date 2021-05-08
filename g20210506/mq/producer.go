package mq

import (
	"sync"
	"sync/atomic"
	"time"
)

type mockedProducer struct {
	total int32
	count int32
	// 使用waitgroup模拟任务的完成
	wait sync.WaitGroup
}

// 实现producer interface 的方法：Produce()
func (p *mockedProducer) Produce() (string, bool) {
	if atomic.AddInt32(&p.count, 1) <= p.total {
		p.wait.Done()
		return "item", true
	}
	time.Sleep(time.Second)
	return "", false
}
