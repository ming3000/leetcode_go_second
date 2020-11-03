package plusplus

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	dfsCombinationSum2(candidates, target, 0, nil, &ret)
	return ret
}

func dfsCombinationSum2(nums []int, target, start int, path []int, ret *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		if nums[i] > target {
			return
		} // if>>
		if i > start && nums[i] == nums[i-1] {
			continue
		} // if>>
		path = append(path, nums[i])
		dfsCombinationSum2(nums, target-nums[i], i+1, path, ret)
		path = path[:len(path)-1]
	} // for>
}
