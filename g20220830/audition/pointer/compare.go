package pointer

import "fmt"

// @see https://topgoer.cn/docs/interview/interview-1dks8a8s7sevg
// Go中两个Nil可能不相等吗？
func compare() {
	var p *int = nil
	var i interface{} = p

	fmt.Println(p == nil)
	fmt.Println(i == nil)
	fmt.Println(p == i)
}
