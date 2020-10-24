package leetcode_go_second

func sumNumbers(root *TreeNode) int {
	var resSum int
	dfsSumNumbers(root, 0, &resSum)
	return resSum
}

func dfsSumNumbers(root *TreeNode, pathSum int, resSum *int) {
	if root == nil {
		return
	}

	pathSum = pathSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*resSum = *resSum + pathSum
		pathSum = 0
		return
	}
	dfsSumNumbers(root.Left, pathSum, resSum)
	dfsSumNumbers(root.Right, pathSum, resSum)
}
