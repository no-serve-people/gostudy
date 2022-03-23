package codetop

// @see https://leetcode-cn.com/problems/sort-an-array/

// mergeSort 返回一个有序切片
func mergeSort(arrs []int) []int {
	n := len(arrs)
	// 当切片只有一个元素的时候，那该切片就是有序的
	if n < 2 {
		return arrs
	}
	// 继续分治
	mid := n / 2
	left := arrs[:mid]
	right := arrs[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

// merge 将两个有序的数组合并成一个有序数组
func merge(left []int, right []int) []int {
	res := make([]int, 0)
	L_left := len(left)
	L_right := len(right)
	i, j := 0, 0
	// 先取二者最优
	for i < L_left && j < L_right {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	// 补上剩下的尾巴
	if i < L_left {
		res = append(res, left[i:]...)
	}
	if j < L_right {
		res = append(res, right[j:]...)
	}

	return res
}
