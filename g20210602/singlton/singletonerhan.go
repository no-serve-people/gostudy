package main

type singleton struct {
}

var instance *singleton

// GetInstance 饿汉模式
// 不加锁
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

