package first

func isValid(s string) bool {
	if s == "" {
		return true
	}
	cm := map[byte]byte{')': '(', ']': '[', '}': '{'}
	cs := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			cs = append(cs, s[i])
		} else {
			if len(cs) == 0 {
				return false
			} // if>>>
			if cm[s[i]] != cs[len(cs)-1] {
				return false
			} // if>>>
			cs = cs[:len(cs)-1]
		} // else>>
	} // for>

	return len(cs) == 0
}
