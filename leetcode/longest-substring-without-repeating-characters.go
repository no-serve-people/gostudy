package leetcode

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	left, right, maxlength := 0, 0, 1

	window := make(map[byte]bool)
	for right < n {
		rightChar := s[right]
		for window[rightChar] {
			delete(window, s[left])
			left++
		}

		if right-left+1 > maxlength {
			maxlength = right - left + 1
		}
		window[rightChar] = true
		right++
	}

	return maxlength
}
