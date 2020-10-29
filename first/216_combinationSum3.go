package first

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	dfsCombinationSum3(k, n, 1, nil, &res)
	return res
}

func dfsCombinationSum3(size, target, start int, path []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 && size == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= 9; i++ {
		path = append(path, i)
		dfsCombinationSum3(size-1, target-i, i+1, path, res)
		path = path[:len(path)-1]
	}
}
