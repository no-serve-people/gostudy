package mq

import (
	"sync/atomic"
)

type mockedConsumer struct {
	count int32
}

func (c *mockedConsumer) Consume(string) error {
	atomic.AddInt32(&c.count, 1)
	return nil
}

//func TestQueue(t *testing.T) {
//	consumer := mockedConsumer{}()
//}
