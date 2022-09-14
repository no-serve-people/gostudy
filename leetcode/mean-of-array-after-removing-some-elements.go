package leetcode

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)

	ans, n := 0, len(arr)/20

	for i := n; i < n*19; i++ {
		ans += arr[i]
	}

	return float64(ans) / float64(n*18)
}
