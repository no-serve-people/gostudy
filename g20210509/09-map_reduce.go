package main

// https://geekr.dev/posts/go-func-map-reduce-filter-mode
// Go 函数式编程篇（六）：引入 Map-Reduce-Filter 模式处理集合元素
// 从处理集合元素聊起
import (
	"fmt"
	"strconv"
)

func ageSum(users []map[string]string) int {
	var sum int
	for _, user := range users {
		num, _ := strconv.Atoi(user["age"])
		sum += num
	}
	return sum
}

func main() {
	users := []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
	}

	fmt.Printf("用户年龄累加结果是:%d\n", ageSum(users))
}
