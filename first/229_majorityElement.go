package first

func majorityElement(nums []int) []int {
	countMap := make(map[int]int, len(nums))
	for _, v := range nums {
		countMap[v]++
	} // for>

	ret := make([]int, 0, len(countMap))
	threshold := len(nums) / 3
	for i, v := range countMap {
		if v > threshold {
			ret = append(ret, i)
		} // if>>
	} // for>
	return ret
}
