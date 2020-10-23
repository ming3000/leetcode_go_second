package leetcode_go_second

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		} // else>>
	} // for>
	return append([]int{1}, digits...)
}
