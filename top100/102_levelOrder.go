package middle

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevelLen := len(queue)
		curLevelList := make([]int, 0)
		for i := 0; i < curLevelLen; i++ {
			head := queue[0]
			queue = queue[1:]
			curLevelList = append(curLevelList, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			} // if>>
			if head.Right != nil {
				queue = append(queue, head.Right)
			} // if>>
		} // for>>

		ret = append(ret, curLevelList)
	} // for>

	return ret
}
