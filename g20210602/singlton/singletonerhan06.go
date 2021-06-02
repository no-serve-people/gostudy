package main

import "sync"

type singleton05 struct {
}

var instance05 *singleton05
var once sync.Once

// GetInstance05 原子操作实现
func GetInstance05() *singleton05 {
	once.Do(func() {
		instance05 = new(singleton05)
	})
	return instance05
}
