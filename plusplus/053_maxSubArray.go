package plusplus

// 动态规划的难点在于找到状态转移方程，
// dp[i] - 表示到当前位置 i 的最大子序列和
// 状态转移方程为： dp[i] = max(dp[i - 1] + nums[i], nums[i])
// 初始化：dp[0] = nums[0]
func maxSubArray(nums []int) int {
	var ret = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		} // if>
		ret = maxInt(ret, nums[i])
	} // for>
	return ret
}
