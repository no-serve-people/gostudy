package main

import "fmt"

// 面试官：go中for-range使用过吗？这几个问题你能解释一下原因吗？
// https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247483875&idx=1&sn=c5da523333e807dddb70228abc2a05b9&scene=21#wechat_redirect
func main() {
	addTomap := func() {
		t := map[string]string{
			"asong":  "太帅",
			"song":   "好帅",
			"asong1": "非常帅",
		}
		for k := range t {
			t["song2020"] = "真话"
			fmt.Printf("%s%s ", k, t[k])
		}
	}

	for i := 0; i < 10; i++ {
		addTomap()
		fmt.Println()
	}
}
