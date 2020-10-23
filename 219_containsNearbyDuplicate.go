package leetcode_go_second

func containsNearbyDuplicate(nums []int, k int) bool {
	checkMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := checkMap[nums[i]]; ok {
			return true
		} // if>>
		checkMap[nums[i]] = true
		if len(checkMap) > k {
			delete(checkMap, nums[i-k])
		}
	} // for>
	return false
}
