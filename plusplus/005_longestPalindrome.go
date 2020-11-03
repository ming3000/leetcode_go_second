package plusplus

// 我们可以用 dp[i][j] 表示 s 中从 i 到 j（包括 i 和 j）是否可以形成回文，
// 状态转移方程只是将上面的描述转化为代码即可：
// if (s[i] === s[j] && dp[i + 1][j - 1]) {
//     dp[i][j] = true;
// }
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var ret string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			} // else>>>
			if dp[i][j] && j-i+1 > len(ret) {
				ret = s[i : j+1]
			}
		} // for>>
	} // for

	return ret
}
