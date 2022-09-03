package leetcode

import (
	"math"
	"sort"
)

// 贪心算法
func findLongestChain(pairs [][]int) (ans int) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	current := math.MinInt32
	for _, pair := range pairs {
		if current < pair[0] {
			current = pair[1]
			ans++
		}
	}
	return ans
}
