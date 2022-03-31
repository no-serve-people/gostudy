package workpool

import "fmt"

type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) error
}

func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{
		Data: data,
		f:    f,
	}
}

// Go语言的并发与WorkerPool - 第二部分
// https://mp.weixin.qq.com/s/u3h8mfNja_SAnFP4vawCGg
func process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	task.Err = task.f(task.Data)
}
