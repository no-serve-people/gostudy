package leetcode

// 三种情况 l1 >=l2 l1 <l2 l1+l2 超出位置了
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head1 := l1
	head2 := l2
	// 先累加
	for head1 != nil {
		if head2 != nil {
			head1.Val += head2.Val
			head2 = head2.Next
		}
		if head1.Next == nil && head2 != nil {
			head1.Next = head2
			// head2 = head2.Next
			break
		}
		head1 = head1.Next
	}
	// 再进位
	mergeAdd(l1)
	return l1
}

func mergeAdd(node *ListNode) {
	for node != nil {
		if node.Val >= 10 {
			node.Val = node.Val % 10
			// 如果是最后一位 补node
			if node.Next == nil {
				node.Next = new(ListNode)
			}
			node.Next.Val += 1
		}
		node = node.Next
	}
}
