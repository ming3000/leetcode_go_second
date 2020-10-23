package leetcode_go_second

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}

	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	res := make([][]int, 0)
	dfsCombine(nums, k, 0, nil, &res)
	return res
}

func dfsCombine(nums []int, k int, start int, path []int, res *[][]int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfsCombine(nums, k, i+1, path, res)
		path = path[:len(path)-1]
	}
}
