package middle

import "container/heap"

// 大顶堆
type BigTopList []int

func (b BigTopList) Len() int {
	return len(b)
}

func (b BigTopList) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BigTopList) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *BigTopList) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *BigTopList) Pop() interface{} {
	ret := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return ret
}

func findKthLargest(nums []int, k int) int {
	temp := make(BigTopList, 0, k)
	var size int
	for _, v := range nums {
		if size < k {
			heap.Push(&temp, v)
			size++
		} else {
			if v > temp[0] {
				heap.Pop(&temp)
				heap.Push(&temp, v)
			} // if>>>
		} // else>>
	} // for>

	return temp[0]
}
