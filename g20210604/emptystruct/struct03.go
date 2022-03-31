package main

import "fmt"

type Set map[string]struct{}

func (s Set) Append(k string) {
	s[k] = struct{}{}
}

func (s Set) Remove(k string) {
	delete(s, k)
}

func (s Set) Exist(k string) bool {
	_, ok := s[k]
	return ok
}

// https://mp.weixin.qq.com/s/8NJLe_11uvu22HfDjKqKjQ
// 详解 Go 空结构体的 3 种使用场景
func main() {
	set := Set{}
	set.Append("煎鱼")
	set.Append("咸鱼")
	set.Append("蒸鱼")
	set.Remove("煎鱼")

	fmt.Println(set.Exist("煎鱼"))
}
