package _struct

import "fmt"

type Door struct{}

func (d Door) Open() {
	fmt.Println("OPNE")
}

func (d Door) Close() {
	fmt.Println("Close")
}
