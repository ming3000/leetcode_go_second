package leetcode_go_second

func searchRange(nums []int, target int) []int {
	last, right := findFirst(nums, target), findLast(nums, target)
	return []int{last, right}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		} // else>>
	} // for>
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		} // else>>
	} // for>
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
