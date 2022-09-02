package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	tmp     = 0
	max     = 0
	leftTmp = 0
)

// 递归法
func longestUnivaluePath(root *TreeNode) int {
	max = 0
	fun(root)
	return max
}

func fun(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Val == root.Val {
		getLen(root.Left, 1)
		leftTmp = tmp
		tmp = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		getLen(root.Right, 1)
	}

	if max < leftTmp+tmp {
		max = leftTmp + tmp
	}

	leftTmp = 0
	tmp = 0
	if root.Left != nil {
		fun(root.Left)
	}

	if root.Right != nil {
		fun(root.Right)
	}
}

func getLen(root *TreeNode, culen int) {
	if culen > tmp {
		tmp = culen
	}

	if root.Left != nil && root.Left.Val == root.Val {
		getLen(root.Left, culen+1)
	}

	if root.Right != nil && root.Right.Val == root.Val {
		getLen(root.Right, culen+1)
	}
}
