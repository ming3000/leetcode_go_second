package first

func removeDuplicates080(nums []int) int {
	j, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count += 1
		} else {
			count = 1
		} // else>>

		if count <= 2 {
			nums[j] = nums[i]
			j += 1
		} // if>>
	} // for>
	return j
}
