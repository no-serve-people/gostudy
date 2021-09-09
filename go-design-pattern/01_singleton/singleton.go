package singleton

// SingleTon 饿汉式单例
type SingleTon struct{}

var singleton *SingleTon

func init() {
	singleton = &SingleTon{}
}

// GetInstance 获取实例
func GetInstance() *SingleTon {
	return singleton
}
