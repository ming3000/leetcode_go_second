package plusplus

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	bucket := make(map[string][]string)
	for _, s := range strs {
		k := generateKey(s)
		bucket[k] = append(bucket[k], s)
	} // for>

	ret := make([][]string, 0)
	for _, v := range bucket {
		ret = append(ret, v)
	}
	return ret
}

func generateKey(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
