package middle

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevel := make([]int, 0)
		curLevelLen := len(queue)
		for i := 0; i < curLevelLen; i++ {
			head := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			} // if>>>
			if head.Right != nil {
				queue = append(queue, head.Right)
			} // if>>>
		} // for>>
		ret = append(ret, curLevel)
	} // for>
	return ret
}
