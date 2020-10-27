package middle

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] > last[1] {
			merged = append(merged, intervals[i])
		} else {
			if last[1] < intervals[i][1] {
				last[1] = intervals[i][1]
			} // if>>>
		} // else>>
	} // for>
	return merged
}
