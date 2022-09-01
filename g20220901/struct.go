package g20220901

import "fmt"

type stPeople struct {
	Gender bool
	Name   string
}

type stStudent struct {
	stPeople
	Class int
}

// 尝试4 嵌套结构的初始化表达式
// var s1 = stStudent{false, "jje", 1}
var s1 = stStudent{stPeople{false, "jje"}, 1}

func testStruct() {
	fmt.Println(s1.Class, s1.Name, s1.Gender)
}
