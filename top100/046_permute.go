package middle

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	usedNums := make(map[int]bool, len(nums))
	dfsPermute(nums, usedNums, nil, &ret)
	return ret
}

func dfsPermute(nums []int, usedNums map[int]bool, path []int, ret *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	for _, num := range nums {
		if usedNums[num] {
			continue
		} // if>>
		usedNums[num] = true
		path = append(path, num)
		dfsPermute(nums, usedNums, path, ret)
		path = path[:len(path)-1]
		usedNums[num] = false
	} // for>
}
