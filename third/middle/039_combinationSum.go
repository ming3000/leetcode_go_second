package middle

func combinationSum039(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfsCombinationSum039(candidates, target, 0, nil, &res)
	return res
}

func dfsCombinationSum039(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfsCombinationSum039(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}
