package leetcode_go_third

func binarySearchFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		} // else>>
	} // for>

	if low < len(nums) && nums[low] == target {
		return low
	}
	return -1
}

func binarySearchLast(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		} // else>>
	} // for>
	if high < len(nums) && nums[high] == target {
		return high
	}
	return -1
}
