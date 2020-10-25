package leetcode_go_third

func maxSubArray(nums []int) int {
	var theMax = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		} // if>>
		theMax = maxInt(theMax, nums[i])
	} // for>
	return theMax
}
