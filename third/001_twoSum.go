package leetcode_go_third

func twoSum(nums []int, target int) []int {
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
