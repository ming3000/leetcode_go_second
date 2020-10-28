package middle

func lengthOfLongestSubstring(s string) int {
	cm := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		cm[s[i]] = 0
	} // for>

	var left int
	var ret int
	for i := 0; i < len(s); i++ {
		for cm[s[i]] != 0 {
			cm[s[left]] -= 1
			left++
		} // for>>
		cm[s[i]] += 1
		ret = maxInt(ret, i-left+1)
	} // for>
	return ret
}
