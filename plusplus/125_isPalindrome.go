package plusplus

import "strings"

// 针对这个问题，我们可以使用头尾双指针，
// 如果两个指针的元素不相同，则直接返回 false,
// 如果两个指针的元素相同，我们同时更新头尾指针，循环。 直到头尾指针相遇。
func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isValidCh(s[left]) {
			left++
			continue
		} // if>>
		if !isValidCh(s[right]) {
			right--
			continue
		} // if>>
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		} // if>>

	} // for>
	return true
}

func isValidCh(ch byte) bool {
	if ch >= '0' && ch <= '9' ||
		ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}
