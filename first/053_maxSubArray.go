package first

import "sort"

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		} // if>>
	} // for>

	sort.Ints(nums)
	return nums[len(nums)-1]
}
