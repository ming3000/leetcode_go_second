package leetcode_go_second

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	} // if>

	var ret int
	var isMinus bool
	str := strings.TrimSpace(s)
	for i, v := range str {
		if i == 0 && v == '-' {
			isMinus = true
			continue
		} // if>>
		if i == 0 && v == '+' {
			isMinus = false
			continue
		} // if>>

		if v >= '0' && v <= '9' {
			ret = ret*10 + int(v-'0')
			if isMinus && ret*(-1) <= math.MinInt32 {
				return math.MinInt32
			} // if>>>
			if !isMinus && ret >= math.MaxInt32 {
				return math.MaxInt32
			} // if>>>
		} else {
			// not a num
			break
		} // if>>
	} // for>

	if isMinus {
		ret = -1 * ret
	}
	return ret
}
