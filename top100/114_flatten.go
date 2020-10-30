package middle

func flatten(root *TreeNode) {
	preList := preOrderTraversal(root)
	for i := 1; i < len(preList); i++ {
		pre, cur := preList[i-1], preList[i]
		pre.Left, pre.Right = nil, cur
	}
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
	preList := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			preList = append(preList, node)
			stack = append(stack, node)
			node = node.Left
		} // for>>
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[:len(stack)-1]
	} // for>
	return preList
}

func preOrderTraversalLoop(root *TreeNode) []*TreeNode {
	preList := make([]*TreeNode, 0)
	if root != nil {
		preList = append(preList, root)
		preList = append(preList, preOrderTraversalLoop(root.Left)...)
		preList = append(preList, preOrderTraversalLoop(root.Right)...)
	}
	return preList
}
