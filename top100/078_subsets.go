package middle

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	dfsSubsets(nums, 0, nil, &ret)
	return ret
}

func dfsSubsets(nums []int, start int, path []int, ret *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ret = append(*ret, temp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfsSubsets(nums, i+1, path, ret)
		path = path[:len(path)-1]
	}
}
