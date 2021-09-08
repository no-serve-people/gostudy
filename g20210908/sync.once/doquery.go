package sync_once

import (
	"sync"
)

// [译] Go sync.Once 的妙用
// https://www.pseudoyu.com/zh/2021/09/02/go_concurrency_sync_once/
// 原生缓存方式
type QueryClient struct {
	cache map[string][]byte
	mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) []byte {
	// 检查结果是否已缓存
	c.mutex.Lock()

	if cached, found := c.cache[name]; found {
		c.mutex.Unlock()
		return cached
	}
	c.mutex.Unlock()

	// 如果未缓存则发出请求
	//resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
	// 为简洁起见，省略了错误处理和 resp.Body.Close
	//result, err := ioutil.ReadAll(resp)
	// 将结果存储在缓存中
	//c.mutex.Lock()
	//c.cache[name] = result
	//c.mutex.Unlock()
	//
	//return result
	return nil
}
