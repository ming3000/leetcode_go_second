package plusplus

func twoSum(nums []int, target int) []int {
	cm := make(map[int]int, len(nums))
	for i, v := range nums {
		diff := target - v
		if index, ok := cm[diff]; ok {
			return []int{i, index}
		}
		cm[v] = i
	} // for>
	return []int{-1, -1}
}
