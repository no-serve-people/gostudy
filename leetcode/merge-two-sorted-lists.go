package leetcode

import "sort"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	list := make([]int, 0)

	h1 := list1
	h2 := list2
	for h1 != nil {
		list = append(list, h1.Val)
		h1 = h1.Next
	}

	for h2 != nil {
		list = append(list, h2.Val)
		h2 = h2.Next
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return newListNode(list)
}

// newListNode 生成一个链表
func newListNode(arr []int) *ListNode {
	var head ListNode
	var pre ListNode
	for _, num := range arr {
		node := ListNode{Val: num, Next: nil}
		if head.Next == nil {
			head.Next = &node
		}
		if pre.Next == nil {
			pre.Next = &node
		}else {
			pre.Next.Next = &node
			pre = *pre.Next
		}
	}
	return head.Next
}