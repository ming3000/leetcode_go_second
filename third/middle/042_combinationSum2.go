package middle

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	dfsCombinationSum042(candidates, target, 0, nil, &res)
	return res
}

func dfsCombinationSum042(candidates []int, target int, start int, path []int, res *[][]int) {
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
		if i > start && candidates[i] == candidates[i-1] {
			continue
		} // if>>
		path = append(path, candidates[i])
		dfsCombinationSum042(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	} // for>
}
