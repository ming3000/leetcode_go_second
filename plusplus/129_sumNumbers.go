package plusplus

func sumNumbers(root *TreeNode) int {
	var ret int
	dfsSumNumbers(root, 0, &ret)
	return ret
}

func dfsSumNumbers(root *TreeNode, pathSum int, ret *int) {
	if root == nil {
		return
	}

	pathSum = pathSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ret += pathSum
		pathSum = 0
		return
	}
	dfsSumNumbers(root.Left, pathSum, ret)
	dfsSumNumbers(root.Right, pathSum, ret)
}
