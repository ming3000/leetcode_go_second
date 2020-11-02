package plusplus

// 假如题目空间复杂度有要求，由于数组是有序的，只需要双指针即可。
// 一个left指针，一个right指针， 如果left + right 值 大于target 则 right左移动， 否则left右移，
func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	} // for>

	return nil
}
