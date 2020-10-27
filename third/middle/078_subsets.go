package middle

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	dfsSubsets(nums, 0, nil, &res)
	return res
}

func dfsSubsets(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfsSubsets(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
