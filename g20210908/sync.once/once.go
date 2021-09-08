package sync_once

import (
	"sync"
)

type CacheEntry03 struct {
	data []byte
	once *sync.Once
}

type QueryClient03 struct {
	cache map[string]*CacheEntry
	mutex *sync.Mutex
}

func (c *QueryClient03) DoQuery03(name string) []byte {
	c.mutex.Lock()
	//entry, found := c.cache[name]
	//
	//if !found {
	//	// 如果在缓存中未找到，创建新的 entry
	//	entry = &CacheEntry03{
	//		once: new(sync.Once),
	//	}
	//	c.cache[name] = entry
	//}
	//c.mutex.Unlock()
	//// 现在，当我们调用 .Do 时，如果有一个正在同步进行的操作
	//// 它将一直阻塞，直到完成（并填充 entry.data）
	//// 或者如果操作之前已经完成过一次
	//// 本次调用不会进行操作，也不会阻塞
	//entry.once.do(func() {
	//	resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
	//	// 为简洁起见，省略了错误处理和 resp.Body.Close
	//	entry.data, err = ioutil.ReadAll(resp)
	//})
	//return entry.data
	return nil
}
