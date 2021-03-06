package singleton

import "sync"

var (
	lazySingleton *SingleTon
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *SingleTon {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &SingleTon{}
		})
	}
	return lazySingleton
}
