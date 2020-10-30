package middle

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	mid := len(nums) / 2
	return nums[mid]
}
