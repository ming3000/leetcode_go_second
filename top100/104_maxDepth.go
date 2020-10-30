package middle

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var deep int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		deep++
		curLevelLen := len(queue)
		for i := 0; i < curLevelLen; i++ {
			head := queue[0]
			queue = queue[1:]
			if head.Left != nil {
				queue = append(queue, head.Left)
			} // if>>
			if head.Right != nil {
				queue = append(queue, head.Right)
			} // if>>
		} // for>>
	} // for>
	return deep
}
