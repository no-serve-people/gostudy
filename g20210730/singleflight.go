package singleflight

import "sync"
//Group https://studygolang.com/articles/18835?fr=sidebar
// singleflight
type Group struct {
	mu sync.Mutex
	m  map[string]*Call
}

//Call call代表需要被执行的函数
type Call struct {
	wg  sync.WaitGroup // 用于阻塞这个调用call的其他请求
	val interface{}    // 函数执行后的结果
	err error          // 函数执行后的error
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	if g.m == nil {
		g.m = make(map[string]*Call)
	}
	// 如果获取当前key的函数正在被执行，则阻塞等待执行中的，等待其执行完毕后获取它的执行结果
	if c, ok := g.m[key]; ok {
		g.mu.Lock()
		c.wg.Wait()
		return c.val, c.err
	}
	// 初始化一个call，往map中写后就解
	c := new(Call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	// 执行获取key的函数，并将结果赋值给这个Call
	c.val, c.err = fn()

	c.wg.Done()
	// 重新上锁删除key
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
