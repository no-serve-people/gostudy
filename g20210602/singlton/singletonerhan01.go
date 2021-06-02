package main

import "sync"

type singleton02 struct {
}

var instance02 *singleton02
var lock sync.Mutex

// GetInstance02 饿汉模式
// 整个方法加锁
func GetInstance02() *singleton02 {
	lock.Lock()
	defer lock.Unlock()
	if instance02 == nil {
		instance02 = new(singleton02)
	}

	return instance02
}
