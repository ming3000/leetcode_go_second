package plusplus

import "math"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 0
	sum, ret := 0, math.MaxInt64
	for right < len(nums) {
		sum += nums[right]
		for sum >= s {
			ret = minInt(ret, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	} // if>

	if ret == math.MaxInt64 {
		return 0
	} else {
		return ret
	}
}
