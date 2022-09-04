package leetcode

// 双指针
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		// 计算当前最大面积
		cur := (j - i) * min(height[i], height[j])

		if cur > res {
			res = cur
		}

		// 移动较小的一侧指针
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
