package g20220901

import "fmt"

func testArray() {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)
}

func testArray2() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(&x)
	fmt.Println(x)
}

func testSlice1() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)
	}

	for _, v := range x {
		fmt.Println(v)
	}
}

func testNilSlice2() {
	var slice1 []int
	// slice1[0] = 1

	slice2 := make([]int, 0)

	slice3 := []int{}

	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
	fmt.Println(slice3 == nil)
}
