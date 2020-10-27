package middle

import "math"

func minSubArrayLen(s int, nums []int) int {
	left, right := 0, -1
	sum, ret := 0, math.MaxInt64

	for left < len(nums) {
		if right+1 < len(nums) && sum < s {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		} // else>>

		if sum > s {
			ret = minInt(ret, right-left+1)
		} // if>>
	} // if>

	if ret == math.MaxInt64 {
		return 0
	} else {
		return ret
	}
}
