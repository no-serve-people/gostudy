package main

import (
	"fmt"
	"net/http"
	"time"
)

// 使用 Go 每分钟处理百万请求
// https://mp.weixin.qq.com/s/BPCxSYEr6ww2F0b9tqPdSQ
type Payload struct {
	// 传啥不重要

}

func (p *Payload) UpdateToS3() error {
	// 存储逻辑,模拟操作耗时
	time.Sleep(500 * time.Millisecond)
	fmt.Println("上传成功")
	return nil
}
func payloadHandler(w http.ResponseWriter, r *http.Request) {
	// 业务过滤
	// 请求body解析......
	var p Payload
	go p.UpdateToS3()
	w.Write([]byte("操作成功"))
}
func main() {
	http.HandleFunc("/payload", payloadHandler)
}
