package first

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	} // if>

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	} // for>
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	theBegin, theLen := 0, 1
	for j := 1; j < len(s); j++ {
		for i := 0; i < len(s); i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if (j-1)-(i+1)+1 < 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				} // else>>>>
			} // else>>>

			if dp[i][j] && j-i+1 > theLen {
				theBegin = i
				theLen = j - i + 1
			} // if>>>
		} // for>>
	} // for>

	return s[theBegin : theBegin+theLen]
}
