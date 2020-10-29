package first

func permute(nums []int) [][]int {
	usedNumMap := make(map[int]bool)
	res := make([][]int, 0)
	dfsPermute(nums, usedNumMap, nil, &res)
	return res
}

func dfsPermute(nums []int, usedNumMap map[int]bool, path []int, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for _, v := range nums {
		if usedNumMap[v] {
			continue
		} // if>>

		usedNumMap[v] = true
		path = append(path, v)
		dfsPermute(nums, usedNumMap, path, res)
		usedNumMap[v] = false
		path = path[:len(path)-1]
	} // for>
}
