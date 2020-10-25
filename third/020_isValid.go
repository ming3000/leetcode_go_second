package leetcode_go_third

func isValid(s string) bool {
	checkMap := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	checkStack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			checkStack = append(checkStack, s[i])
		} else {
			if len(checkStack) == 0 {
				return false
			} // if>>>

			top := checkStack[len(checkStack)-1]
			checkStack = checkStack[:len(checkStack)-1]
			if s[i] != checkMap[top] {
				return false
			} // if>>>
		} // else>>
	} // for>

	return len(checkStack) == 0
}
