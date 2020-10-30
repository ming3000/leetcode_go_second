package middle

func singleNumber(nums []int) int {
	var ret = 0
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}
