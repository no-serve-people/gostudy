package workerpool

import "sync"

// https://github.com/lgphone/workerpool/blob/main/workerpool.go
// go 协程并发控制，协程错误接收
type workerPool struct {
	mutex sync.Mutex
	ch    chan bool
	errs  []error
}

func (s *workerPool) Submit(fn func()) {
	s.ch <- true
	go func() {
		defer func() {
			<-s.ch
		}()

		defer func() {
			if err := recover(); err != nil {
				s.mutex.Lock()
				s.errs = append(s.errs, err.(error))
				s.mutex.Unlock()
			}
		}()
		fn()
	}()
}

func (s *workerPool) Wait() []error {
	for {
		if len(s.ch) == 0 {
			close(s.ch)
			return s.errs
		}
	}
}

func NewWorkerPool(maxWorkerNumber int) *workerPool {
	return &workerPool{
		ch:   make(chan bool, maxWorkerNumber),
		errs: make([]error, 0),
	}
}
