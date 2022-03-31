package main

// go解锁设计模式之单例模式
// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247483947&idx=2&sn=32442f6f855162c2185a10e4329c74fa&scene=21#wechat_redirect
type singleton06 struct{}

var instance06 = new(singleton06)

func GetInstance06() *singleton06 {
	return instance06
}

type singleton07 struct{}

var instance07 *singleton07

func init() {
	instance07 = new(singleton07)
}

func GetInstance07() *singleton07 {
	return instance07
}
