package middle

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTCore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCore(root *TreeNode, theMin, theMax int) bool {
	if root == nil {
		return true
	}

	if root.Val <= theMin || root.Val >= theMax {
		return false
	}

	isLeftValid := isValidBSTCore(root.Left, theMin, root.Val)
	isRightValid := isValidBSTCore(root.Right, root.Val, theMax)
	return isLeftValid && isRightValid
}
