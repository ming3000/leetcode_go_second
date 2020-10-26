package middle

func lengthOfLongestSubstring(s string) int {
	checkMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		checkMap[s[i]] = 0
	}

	var left int
	var ret int
	for i := 0; i < len(s); i++ {
		for checkMap[s[i]] != 0 {
			checkMap[s[left]] -= 1
			left += 1
		} // for>>
		checkMap[s[i]] += 1
		ret = maxInt(ret, i-left+1)
	} // for>

	return ret
}
