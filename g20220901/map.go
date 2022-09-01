package g20220901

import "fmt"

func testMap() {
	// 错误的 key 检测方式
	x := map[string]string{"one": "2", "two": "", "three": "3"}

	if v := x["two"]; v == "" {
		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}
}
