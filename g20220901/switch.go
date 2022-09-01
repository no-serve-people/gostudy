package g20220901

import "fmt"

func testSwitch() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ':
			// 空格符会直接 break，返回 false // 和其他语言不一样
			// fallthrough
		case '\t':
			return true

		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}
