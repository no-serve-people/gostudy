package g20220901

import "fmt"

func testNilInterface() {
	var data *byte
	var in interface{}
	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)
	in = data
	fmt.Println(in, in == nil)
}

func testNilInterface2() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil
		}
		return result
	}

	if res := doIt(1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res) // Bad result: <nil>
	}
}
