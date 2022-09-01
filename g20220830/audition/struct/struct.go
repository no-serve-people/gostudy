package _struct

import (
	"fmt"
	"unsafe"
)

func testStruct() {
	fmt.Println("struct", unsafe.Sizeof(struct{}{}))
}
