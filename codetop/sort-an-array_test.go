package codetop

import (
	"fmt"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	numbers := []int{5, 1, 9, 2, 6, 3, 9, 4, 1, 6, 3, 7}
	fmt.Printf("before arr is: %#v\n", numbers)
	res := mergeSort(numbers)
	fmt.Printf("after arr is: %#v\n", res)
}
