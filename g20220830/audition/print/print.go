package print

import "fmt"

type student struct {
	id   int32
	name string
}

func testPrint() {
	a := &student{id: 1, name: "微课鸟窝"}
	fmt.Printf("a=%v \n", a)
	fmt.Printf("a=%+v \n", a)
	fmt.Printf("a=%#v \n", a)
}
