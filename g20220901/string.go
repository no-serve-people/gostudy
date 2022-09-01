package g20220901

import "fmt"

func testString() {
	// x := "text"
	// x[0] = "T"
	// fmt.Println("x")

	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T'
	x = string(xBytes)
	fmt.Println(x) // Text
}
