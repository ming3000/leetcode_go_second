package middle

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 如果大家这样去思考的话，问题就会得到简化。
	// 我们的目标就是分别在左右子树进行查找p和q。
	// 如果p没有在左子树，那么它一定在右子树（题目限定p一定在树中）， 反之亦然。
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
