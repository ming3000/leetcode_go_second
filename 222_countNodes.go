package leetcode_go_second

func countHeight(root *TreeNode) int {
	var height int
	for root != nil {
		root = root.Left
		height += 1
	}
	return height
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := countHeight(root.Left)
	rightHeight := countHeight(root.Right)
	if leftHeight == rightHeight {
		return countNodes(root.Right) + (1 << leftHeight)
	} else {
		return countNodes(root.Right) + (1 << rightHeight)
	}
}
