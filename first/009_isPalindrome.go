package first

func isPalindrome(x int) bool {
	// 负数肯定不是回文，因为负号
	if x < 0 {
		return false
	} // if>
	raw, reversed := x, 0
	for raw != 0 {
		reversed = reversed*10 + raw%10
		raw /= 10
	} // for>

	return x == reversed
}
