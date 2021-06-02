package main

import "sync"

type singleton04 struct {
}

var instance04 *singleton04
var lock04 sync.Mutex

// GetInstance04 饿汉模式
// 双重检锁
func GetInstance04() *singleton04 {
	if instance04 == nil {
		lock04.Lock()
		if instance04 == nil {
			instance04 = new(singleton04)
		}
		lock04.Unlock()
	}

	return instance04
}
