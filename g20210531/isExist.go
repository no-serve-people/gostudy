package main

import "fmt"

// https://juejin.cn/post/6960991873112473614
// 漫画| Golang 二维数组中查找值是否存在
/**
题目：二维数组中查找值是否存在
在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
{
	{1,4,7,11,15},
	{2,5,8,12,19},
	{10,6,9,16,22},
	{20,13,14,17,24},
	{18,21,23,26,30},
}

给定 target = 7，返回 true。
给定 target = 3，返回 false。
*/

/**
方法一：暴力
不考虑数组 每一行已经排好序 的特点，直接遍历二维数组每一行和每一列每个元素，
如果找到元素等于目标值，则返回 true。否则返回 false。
*/
func isExistV2(matrix [][]int, target int) bool {
	i := 0 // matrix[0] 最开始一行数据 {1,2,8,9}
	j := 0 // matrix[0][0] 最开始元素 1
	for i < len(matrix) {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i++ //由于 每一行都按照从左到右递增的顺序排序，target < 该行第一个元素，意味着也小于该行所有元素
				j = 0
			} else if target > matrix[i][j] {
				j++ // target > 该行第一个元素，就继续对比下一个
			} else if target == matrix[i][j] {
				return true
			}
		} else {
			return false //超出数组返回false
		}
	}
	return false //超出matrix返回false
}

func isExistV1(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

func main() {
	temArr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{10, 6, 9, 16, 22},
		{20, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	ret1 := isExistV1(temArr, 7)
	fmt.Println("check 7 is exist:", ret1)

	ret2 := isExistV2(temArr, 3)
	fmt.Println("check 3 is exist:", ret2)
}
