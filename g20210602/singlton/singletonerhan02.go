package main

import "sync"

type singleton03 struct{}

var (
	instance03 *singleton03
	lock02     sync.Mutex
)

// GetInstance03 饿汉模式
// 创建方法时进行锁定
func GetInstance03() *singleton03 {
	if instance03 == nil {
		lock02.Lock()
		instance03 = new(singleton03)
		lock02.Unlock()
	}

	return instance03
}
