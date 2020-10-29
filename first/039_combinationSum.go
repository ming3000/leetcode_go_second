package first

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	dfsCombinationSum(candidates, target, 0, nil, &res)
	return res
}

func dfsCombinationSum(candidates []int, target int, start int, path []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[start] > target {
			return
		}
		path = append(path, candidates[i])
		dfsCombinationSum(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	} // for>
}
