package rune

import "fmt"

func testRune() {
	str := "hello 你好"
	// golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算
	// 所以一个汉字占3个字节算了3个长度
	fmt.Printf("length %d \n", len(str))
	fmt.Printf("length %d \n", len([]byte(str)))
	fmt.Printf("length rune %d \n", len([]rune(str)))
}
