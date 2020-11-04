package plusplus

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTCode(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCode(root *TreeNode, theMin, theMax int) bool {
	if root == nil {
		return true
	}
	if root.Val <= theMin || root.Val >= theMax {
		return false
	}
	leftValid := isValidBSTCode(root.Left, theMin, root.Val)
	rightValid := isValidBSTCode(root.Right, root.Val, theMax)
	return leftValid && rightValid
}
