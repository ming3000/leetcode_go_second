package middle

func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	queue := []*TreeNode{u, v}
	for len(queue) > 0 {
		u, v = queue[0], queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		} // if>>
		if u == nil || v == nil {
			return false
		} // if>>
		if u.Val != v.Val {
			return false
		} // if>>
		queue = append(queue, u.Left)
		queue = append(queue, v.Right)
		queue = append(queue, u.Right)
		queue = append(queue, v.Left)
	} // for>
	return true
}
