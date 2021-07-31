package main

import "log"
// Base 生产环境遇到一个 Go 问题，整组人都懵逼了...
// https://juejin.cn/post/6982038225514659870?from=main_page#comment
type Base interface {
	do()
}
type App struct {
}

func main() {
	var base Base
	base = GetApp()
	log.Println(base)
	log.Println(base == nil)
}

func GetApp() *App {
	return nil
}

func (a *App) do() {}
