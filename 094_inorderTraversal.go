package leetcode_go_second

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	tempStack := make([]*TreeNode, 0)
	for len(tempStack) > 0 || root != nil {
		for root != nil {
			tempStack = append(tempStack, root)
			root = root.Left
		} // for>>
		ret = append(ret, tempStack[len(tempStack)-1].Val)
		root = tempStack[len(tempStack)-1].Right
		tempStack = tempStack[:len(tempStack)-1]
	} // for>
	return ret
}
