package leetcode

func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return
}
