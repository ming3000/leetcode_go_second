package easy

import "strings"

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	bArr := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isValidChar(s[i]) {
			bArr = append(bArr, s[i])
		} // if>>
	} // for>

	for i := 0; i < len(bArr)/2; i++ {
		if bArr[i] != bArr[len(bArr)-1-i] {
			return false
		} // if>>
	} // for>
	return true
}

func isValidChar(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
