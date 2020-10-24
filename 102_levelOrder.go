package leetcode_go_second

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLevelLen := len(queue)
		curLevelRet := make([]int, 0, curLevelLen)
		for i := 0; i < curLevelLen; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevelRet = append(curLevelRet, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			} // if>>>
			if node.Right != nil {
				queue = append(queue, node.Right)
			} // if>>>
		} // for>>
		ret = append(ret, curLevelRet)
	} // for>
	return ret
}
