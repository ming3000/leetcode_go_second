package leetcode_go_second

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else {
			return cur
		} // else>>
	} // for>

	return nil
}
