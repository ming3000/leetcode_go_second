package first

import "leetcode_go_second"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	var ret int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = leetcode_go_second.maxInt(dp[i], dp[j]+1)
			} // if>>>
		} // for>>
		ret = leetcode_go_second.maxInt(ret, dp[i])
	} // for>
	return ret
}
