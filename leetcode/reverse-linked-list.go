package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 一、最简单的思路：修改所有指针方向即可
	//// 二、需要几个变量：两个？还是三个？
	//// 三、思考边界条件
	//// 以上三个顺序不能乱，否则的话解题效率很低
	//// (四、循环怎么选择？ for? each?)
	//// (五、循环条件？)
	////var pre *ListNode
	////cur := head
	////for cur != nil {
	////	next := cur.Next // 工具人变量，完全是为了移动cur用的
	////	cur.Next = pre   // 关键步骤：修改指针方向
	////	pre = cur
	////	cur = next
	////}
	////return pre
	//cur := head
	//var pre *ListNode = nil
	//for cur != nil {
	//	pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	//}
	//return pre
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 工具人变量，完全是为了移动cur用的
		cur.Next = pre   // 关键步骤：修改指针方向
		pre = cur
		cur = next
	}

	return pre
}
