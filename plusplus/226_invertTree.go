package plusplus

// 算法：
// 遍历树（随便怎么遍历），然后将左右子树交换位置。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
