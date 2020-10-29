package first

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	dfsSubsetsWithDup(nums, 0, nil, &res)
	return res
}

func dfsSubsetsWithDup(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		} // if>>
		path = append(path, nums[i])
		dfsSubsetsWithDup(nums, i+1, path, res)
		path = path[:len(path)-1]
	} // for>
}
