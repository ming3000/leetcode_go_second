package middle

func sumNumbers(root *TreeNode) int {
	var res int
	dfsSumNumbers(root, 0, &res)
	return res
}

func dfsSumNumbers(root *TreeNode, pathSum int, res *int) {
	if root == nil {
		return
	}

	pathSum = pathSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res = *res + pathSum
		pathSum = 0
		return
	}
	dfsSumNumbers(root.Left, pathSum, res)
	dfsSumNumbers(root.Right, pathSum, res)
}
