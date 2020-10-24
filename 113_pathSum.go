package leetcode_go_second

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	dfsPathSum(root, sum, nil, &res)
	return res
}

func dfsPathSum(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			temp = append(temp, root.Val)
			*res = append(*res, temp)
		}
		return
	}
	path = append(path, root.Val)
	dfsPathSum(root.Left, sum-root.Val, path, res)
	dfsPathSum(root.Right, sum-root.Val, path, res)
	path = path[:len(path)-1]
}
