package plusplus

func permute(nums []int) [][]int {
	usedNums := make(map[int]bool)
	ret := make([][]int, 0)
	dfsPermute(nums, usedNums, nil, &ret)
	return ret
}

func dfsPermute(nums []int, usedNums map[int]bool, path []int, ret *[][]int) {
	if len(nums) == len(path) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}

	for _, v := range nums {
		if usedNums[v] {
			continue
		} // if>>
		usedNums[v] = true
		path = append(path, v)
		dfsPermute(nums, usedNums, path, ret)
		path = path[:len(path)-1]
		usedNums[v] = false
	} // for>
}
