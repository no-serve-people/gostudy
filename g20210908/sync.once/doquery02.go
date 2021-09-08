package sync_once

import (
	"sync"
)

type CacheEntry struct {
	data []byte
	wait <-chan struct{}
}

type QueryClient02 struct {
	cache map[string]*CacheEntry
	mutex *sync.Mutex
}

func (c *QueryClient02) DoQuery02(name string) []byte {
	// 检查操作是否已启动
	c.mutex.Lock()
	if cached, found := c.cache[name]; found {
		c.mutex.Unlock()
		// 等待完成
		<-cached.wait
		return cached.data
	}

	//entry := &CacheEntry{
	//	data: result,
	//	wait: make(chan struct{}),
	//}
	//c.cache[name] = entry
	//c.mutex.Unlock()
	//// 如果未缓存，则发出请求
	//resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
	//// 为简洁起见，省略了错误处理和 resp.Body.Close
	//entry.data, err = ioutil.ReadAll(resp)
	//
	//// 关闭 channel，传递操作完成信号
	//// 立即返回
	//close(entry.wait)
	//
	//return entry.data
	return nil
}
