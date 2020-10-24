package leetcode_go_second

import "strings"

func isPalindrome3(s string) bool {
	s = strings.ToLower(s)
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			b = append(b, s[i])
		} // if>>
	} // for>

	for i := 0; i < len(b)/2; i++ {
		if b[i] != b[len(b)-1-i] {
			return false
		} // if>>
	} // for>
	return true
}

func isalnum(ch byte) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= 'a' && ch < 'z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}
