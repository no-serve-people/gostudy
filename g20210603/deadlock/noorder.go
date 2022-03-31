package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 大佬们求解一个 go map 无序的问题
// https://www.v2ex.com/t/780962#reply26
func main() {
	jsonStr := `{"name":"tom","user_id":"123"}` // 这是传入的参数，name 与 user_id 顺序不能确定先后
	var str string

	m := make(map[string]interface{})

	_ = json.Unmarshal([]byte(jsonStr), &m)
	v := reflect.ValueOf(m)
	keys := v.MapKeys()

	for _, key := range keys {
		v1 := v.MapIndex(key).Interface().(string)
		str += v1
	}
	fmt.Println(str)
	// 由于 map 无序，不能固定输出：tom123
	// 如何保持与 json 中键一致，固定输出？
	// 比如若 json_str := `{"user_id":"123""name":"tom"}` 则输出 123tom
}
