package first

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(floatHeight(root.Left)-floatHeight(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func floatHeight(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(floatHeight(node.Left), floatHeight(node.Right)) + 1
}
