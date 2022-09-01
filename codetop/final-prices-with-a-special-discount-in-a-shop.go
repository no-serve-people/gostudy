package codetop

// 输入：prices = [8,4,6,2,3]
// 输出：[4,2,4,2,3]
// 输入：prices = [1,2,3,4,5]
// 输出：[1,2,3,4,5]
// 输入：prices = [10,1,1,6]
// 输出：[9,0,1,6]
func finalPrices(prices []int) []int {
	n := len(prices)
	if n < 2 {
		return prices
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		discount := 0
		j := i + 1
		for j < n {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
			j++
		}

		res[i] = prices[i] - discount
	}

	return res
}
