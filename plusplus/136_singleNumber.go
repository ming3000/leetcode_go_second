package plusplus

func singleNumber(nums []int) int {
	var ret int
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}
