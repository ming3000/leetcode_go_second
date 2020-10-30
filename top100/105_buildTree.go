package middle

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] {
			break
		} // if>>
	} // for>

	root := &TreeNode{Val: preorder[0]}
	preLeft := preorder[1 : len(inorder[:rootIndex])+1]
	preRight := preorder[len(inorder[:rootIndex])+1:]
	inLeft := inorder[:rootIndex]
	inRight := inorder[rootIndex+1:]
	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, inRight)
	return root
}
