package plusplus

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	dfsSubsetsWithDup(nums, 0, nil, &ret)
	return ret
}

func dfsSubsetsWithDup(nums []int, start int, path []int, ret *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ret = append(*ret, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		} // if>>
		path = append(path, nums[i])
		dfsSubsetsWithDup(nums, i+1, path, ret)
		path = path[:len(path)-1]
	} // for>
}
