package middle

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	dfsCombinationSum(candidates, target, 0, nil, &ret)
	return ret
}

func dfsCombinationSum(candidates []int, target, start int, path []int, ret *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfsCombinationSum(candidates, target-candidates[i], i, path, ret)
		path = path[:len(path)-1]
	}
}
