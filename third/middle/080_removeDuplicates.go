package middle

func removeDuplicates(nums []int) int {
	var slow = 0
	var k = 2
	for fast := 0; fast < len(nums); fast++ {
		if slow < k || nums[fast] != nums[slow-k] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
