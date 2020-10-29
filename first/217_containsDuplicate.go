package first

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		} // if>>
	} // for>
	return false
}
