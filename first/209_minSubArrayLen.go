package first

import (
	"leetcode_go_second"
	"math"
)

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

		if sum >= s {
			ret = leetcode_go_second.minInt(ret, right-left+1)
		} // if>>
	} // for>

	if ret == math.MaxInt64 {
		return 0
	} else {
		return ret
	}
}
