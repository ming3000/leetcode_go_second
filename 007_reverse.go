package leetcode_go_second

import "math"

func reverse(x int) int {
	raw, reversed := x, 0
	for raw != 0 {
		reversed = reversed*10 + raw%10
		raw = raw / 10
		if reversed <= math.MinInt32 || reversed >= math.MaxInt32 {
			return 0
		} // if>>
	} // for>
	return reversed
}
