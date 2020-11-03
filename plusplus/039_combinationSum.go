package plusplus

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	dfsCombinationSum(candidates, target, 0, nil, &ret)
	return ret
}

func dfsCombinationSum(nums []int, target, start int, path []int, ret *[][]int) {
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
		path = append(path, nums[i])
		dfsCombinationSum(nums, target-nums[i], i, path, ret)
		path = path[:len(path)-1]
	}
}
