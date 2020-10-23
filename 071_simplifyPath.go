package leetcode_go_second

import "strings"

func simplifyPath(path string) string {
	itemStack := make([]string, 0, len(path))
	items := strings.Split(path, "/")
	for _, v := range items {
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if len(itemStack) > 0 {
				itemStack = itemStack[:len(itemStack)-1]
			} // if>>>
		} else {
			itemStack = append(itemStack, v)
		} // else>>
	} // for>

	return "/" + strings.Join(itemStack, "/")
}
