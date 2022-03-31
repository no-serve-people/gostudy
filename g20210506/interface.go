package main

import "fmt"

// https://juejin.cn/post/6945476531439271972 聊一聊 Go 的接口 | Go主题月
type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s,code=%d", e.Message, e.Code)
}

type Cat struct{}

type Duck interface{}

// func (c Cat) Quack{}

////func (c *Cat) Quack {
////
////}
//
//var d Duck = Cat{}
//var d Duck = &Cat{}
