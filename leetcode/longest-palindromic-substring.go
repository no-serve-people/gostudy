package leetcode

var srange = [2]int{}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	for i := 1; i < n; i++ {
		helper(s, n, i, i+1)
		helper(s, n, i, i)
	}

	return s[srange[0]:srange[1]]
}

func helper(s string, n, start, end int) {
	for start >= 0 && end <= n-1 {
		if s[start] == s[end] {
			start--
			end++
		} else {
			break
		}
	}

	if end-(start+1) > srange[1]-srange[0] {
		srange[0] = start + 1
		srange[1] = end
	}
}
