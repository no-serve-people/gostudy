package main

import (
	"log"
	"net/http"
)

// 使用 Go 每分钟处理百万请求
// https://mp.weixin.qq.com/s/BPCxSYEr6ww2F0b9tqPdSQ
const MaxQueue = 400

var Queue chan Payload

func init() {
	Queue = make(chan Payload, MaxQueue)
}

func payloadHandler02(w http.ResponseWriter, r *http.Request) {
	// 业务过滤
	// 请求body解析......
	var p Payload
	// go p.UpdateToS3()
	Queue <- p
	w.Write([]byte("操作成功"))
}

// StartProcessor 处理任务
func StartProcessor() {
	for {
		select {
		case payload := <-Queue:
			payload.UpdateToS3()
		}
	}
}

func main() {
	http.HandleFunc("/payload", payloadHandler02)
	// 单独开一个g接收与处理任务
	go StartProcessor()
	log.Fatal(http.ListenAndServe(":8099", nil))
}
