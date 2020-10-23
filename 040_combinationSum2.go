package leetcode_go_second

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfsCombinationSum2(candidates, target, 0, nil, &res)
	return res
}

func dfsCombinationSum2(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		} // if>>
		if i != start && candidates[i] == candidates[i-1] {
			continue
		} // if>>
		path = append(path, candidates[i])
		dfsCombinationSum2(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	} // for>
}
