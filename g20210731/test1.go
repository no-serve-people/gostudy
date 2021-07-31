package main

import "log"

type MyErr struct {
	Msg string
}
// Base 生产环境遇到一个 Go 问题，整组人都懵逼了...
// https://juejin.cn/post/6982038225514659870?from=main_page#comment
func main() {
	var e error
	e = GetErr()
	log.Println(e == nil)
}

func GetErr() *MyErr {
	return nil
}

func (m *MyErr) Error() string {
	return "脑子进煎鱼了"
}
