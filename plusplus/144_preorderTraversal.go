package plusplus

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		} // if>>
		if node.Left != nil {
			stack = append(stack, node.Left)
		} // if>>
	} // for>

	return ret
}
