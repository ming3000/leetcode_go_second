package first

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var height int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		height += 1
		curLevelLen := len(queue)
		for i := 0; i < curLevelLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			} // if>>>
			if node.Right != nil {
				queue = append(queue, node.Right)
			} // if>>>
		} // for>>
	} // for>
	return height
}
