package middle

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	pathSumCore(root, sum, nil, &res)
	return res
}

func pathSumCore(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			temp := make([]int, len(path))
			copy(temp, path)
			temp = append(temp, root.Val)
			*res = append(*res, temp)
		}
		return
	}

	path = append(path, root.Val)
	pathSumCore(root.Left, sum-root.Val, path, res)
	pathSumCore(root.Right, sum-root.Val, path, res)
	path = path[:len(path)-1]
}
