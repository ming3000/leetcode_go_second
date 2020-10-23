package leetcode_go_second

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	checkArr := [26]int{}
	for _, v := range s {
		checkArr[v-'a'] += 1
	}
	for _, v := range t {
		checkArr[v-'a'] -= 1
	}
	for _, v := range checkArr {
		if v != 0 {
			return false
		} // if>>
	} // for>
	return true
}
