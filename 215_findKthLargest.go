package leetcode_go_second

import "container/heap"

type TopList []int

func (t TopList) Len() int {
	return len(t)
}

func (t TopList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TopList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t *TopList) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *TopList) Pop() interface{} {
	ret := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return ret
}

func findKthLargest(nums []int, k int) int {
	temp := make(TopList, 0, k)
	sizeCount := 0
	for _, v := range nums {
		if sizeCount < k {
			heap.Push(&temp, v)
			sizeCount++
		} else {
			if v > temp[0] {
				heap.Pop(&temp)
				heap.Push(&temp, v)
			}
		} // else>>
	} // for>

	return temp[0]
}
