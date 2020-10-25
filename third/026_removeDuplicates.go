package leetcode_go_third

func removeDuplicates(nums []int) int {
	var slow int
	for i := 0; i < len(nums); i++ {
		if nums[slow] != nums[i] {
			slow++
			nums[slow] = nums[i]
		}
	} // for>
	return slow + 1
}
