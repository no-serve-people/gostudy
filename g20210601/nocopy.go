package main

import "fmt"

// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247484551&idx=1&sn=aff008de8b6c089d7daf9c963ce8b60a&scene=21#wechat_redirect
// 源码剖析sync.WaitGroup(文末思考题你能解释一下吗?)

type User struct {
	Name string
	Info *Info
}

type Info struct {
	Age    int
	Number int
}

func main() {
	u := User{
		Name: "asong",
		Info: &Info{
			Age:    10,
			Number: 24,
		},
	}

	u1 := u
	u1.Name = "Golang梦工厂"
	u1.Info.Age = 30
	fmt.Println(u.Info.Age, u.Name)
	fmt.Println(u.Info.Age, u1.Name)
}
