package leetcode_go_second

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTCore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCore(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}
	isLeftValid := isValidBSTCore(root.Left, minVal, root.Val)
	isRightValid := isValidBSTCore(root.Right, root.Val, maxVal)
	return isLeftValid && isRightValid
}
