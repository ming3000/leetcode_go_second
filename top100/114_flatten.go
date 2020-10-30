package middle

func flatten(root *TreeNode) {
	preList := preOrderTraversal(root)
	for i := 1; i < len(preList); i++ {
		pre, cur := preList[i-1], preList[i]
		pre.Left, pre.Right = nil, cur
	}
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	preList := make([]*TreeNode, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		preList = append(preList, node)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
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
