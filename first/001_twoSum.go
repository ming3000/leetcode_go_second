package first

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	cm := make(map[int]int, len(nums))
	for i, v := range nums {
		diff := target - v
		if otherIndex, ok := cm[diff]; ok {
			return []int{i, otherIndex}
		} // if>>
		cm[v] = i
	} // for>
	return nil
}
