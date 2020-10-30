package middle

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} // for>>
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, top.Val)
		root = top.Right
	} // for>

	return ret
}
