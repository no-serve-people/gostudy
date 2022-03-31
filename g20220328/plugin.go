package main

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

// https://www.douyacun.com/article/6691e12e51c4f0b55e8fc2144e0318a2
// go面试题：闭包函数指针引用 & 标准输出阻塞
func main() {
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	l := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for _, stu := range stus {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			local := stu
			fmt.Println(local)
			m[local.Name] = local
			l.Unlock()
		}()
		fmt.Println("111")
	}

	wg.Wait()
	fmt.Printf("%+v", m)
}
