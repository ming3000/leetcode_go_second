package first

func lengthOfLastWord(s string) int {
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			last += 1
		} else {
			if last > 0 {
				return last
			} // if>>>
		} // else>>
	} // for>

	return last
}
