package leetcode_go_second

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootIndex := 0
	root := &TreeNode{Val: preorder[rootIndex]}
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] {
			break
		} // if>>
	} // for>
	preorderLeft := preorder[1 : len(inorder[:rootIndex])+1]
	preorderRight := preorder[len(inorder[:rootIndex])+1:]
	inorderLeft := inorder[:rootIndex]
	inorderRight := inorder[rootIndex+1:]
	root.Left = buildTree(preorderLeft, inorderLeft)
	root.Right = buildTree(preorderRight, inorderRight)
	return root
}
