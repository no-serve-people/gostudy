package g20220808

import (
	"fmt"
	"unsafe"
)

type Programmer2 struct {
	name     string
	age      int
	language string
}

func testPointer2() {
	p := Programmer2{
		name:     "stefno",
		age:      18,
		language: "go",
	}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "golang"

	fmt.Println(p)
}
