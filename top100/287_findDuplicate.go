package middle

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid, count := (left+right)/2, 0
		for _, v := range nums {
			if v <= mid {
				count++
			} // if>>>
		} // for>>
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	} // for>
	return left
}
