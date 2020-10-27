package middle

func kthSmallest(root *TreeNode, k int) int {
	var ret int
	stack := []*TreeNode{root}
	for k > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		} // for>>
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
		ret = top.Val
		k--
	} // for>
	return ret
}
