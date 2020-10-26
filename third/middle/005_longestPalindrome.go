package middle

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var res string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		} // for>>
	} // for>

	return res
}
