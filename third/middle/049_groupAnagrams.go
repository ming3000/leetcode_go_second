package middle

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	cm := make(map[string][]string)
	for _, s := range strs {
		k := generateKey(s)
		if _, ok := cm[k]; !ok {
			cm[k] = make([]string, 0)
		} // if>>
		cm[k] = append(cm[k], s)
	} // for>

	ret := make([][]string, 0)
	for k := range cm {
		ret = append(ret, cm[k])
	}
	return ret
}

func generateKey(str string) string {
	str = strings.TrimSpace(str)
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
