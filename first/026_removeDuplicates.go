package first

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	rightIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[rightIndex] = nums[i]
			rightIndex += 1
		}
	} // for>
	retNums := nums[:rightIndex]
	return len(retNums)
}
