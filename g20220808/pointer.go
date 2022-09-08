package g20220808

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func testPointer1() {
	p := Programmer{
		name:     "stefno",
		language: "go",
	}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "golang"

	fmt.Println(p)
}
