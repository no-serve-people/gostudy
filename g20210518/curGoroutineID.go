package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// https://mp.weixin.qq.com/s/pY_QKi6fuoPBrqo0typkDA
// 进程、线程都有 ID，为什么 Goroutine 没有 ID？
func main() {
	bp := littleBuf.Get().(*[]byte)
	defer littleBuf.Put(bp)
	b := *bp
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	fmt.Println(n)
}

var littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}

var goroutineSpace = []byte("goroutine ")
