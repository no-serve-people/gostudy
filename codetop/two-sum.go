package codetop

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
// 栈操作
func isValid(str string) bool {
	var stack []string
	for _, item := range str {
		s := string(item)
		if s == "{" || s == "[" || s == "(" {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (top == "{" && s == "}") || (top == "[" && s == "]") || (top == "(" && s == ")") {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
