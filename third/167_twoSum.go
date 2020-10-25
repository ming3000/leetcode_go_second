package leetcode_go_third

func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		} // else>>
	} // for>
	return nil
}
