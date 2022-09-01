package g20220901

import "sort"

// IsExist 如何在切片中查找
func IsExist(s []string, t string) (int, bool) {
	iIndex := sort.SearchStrings(s, t)
	bExist := iIndex != len(s) && s[iIndex] == t

	return iIndex, bExist
}
