package codetop

import "strings"

// 两数之和
func twoSum(nums []int, target int) []int {
	var index []int
	for i, num := range nums {
		for j, numj := range nums {
			if num+numj == target && i != j {
				index = append(index, i, j)
				return index
			}
		}

	}

	return nil
}

// 有效的括号
func isValid(s string) bool {
	a := [3][2]string{
		{"(", ")"},
		{"[", "]"},
		{"{", "}"},
	}

	length := len([]rune(s))
	for _, temp := range a {
		index := strings.Index(s, temp[0])
		if index >= 0 {
			for i := 1;; i = i + 2 {
				if i >= length {
					return false
				}
				u := s[index+i]
				if string(u) == temp[1] {
					return true
				}
			}
		}
	}

	return false
}
