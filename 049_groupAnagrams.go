package leetcode_go_second

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	tempMap := make(map[string][]string)
	retSlice := make([][]string, 0)

	for _, v := range strs {
		key := calStrKey(v)
		if _, ok := tempMap[key]; !ok {
			tempMap[key] = make([]string, 0)
		} // if>>
		tempMap[key] = append(tempMap[key], v)
	} // for>

	for k := range tempMap {
		retSlice = append(retSlice, tempMap[k])
	}
	return retSlice
}

func calStrKey(s string) string {
	s = strings.TrimSpace(s)
	a := []byte(s)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return string(a)
}
