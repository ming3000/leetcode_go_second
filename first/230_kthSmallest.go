package first

func kthSmallest(root *TreeNode, k int) int {
	var ret int

	stack := make([]*TreeNode, 0)
	for k > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} // for>>
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = cur.Right
		ret = cur.Val
		k--
	} // for>

	return ret
}
