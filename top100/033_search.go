package middle

func findPivot(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return 0
	}

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		} else {
			if nums[left] < nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		} // else>>
	} // for>
	return 0
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right, pivot := 0, len(nums)-1, findPivot(nums)
	if pivot == 0 || nums[0] > target {
		left, right = pivot, len(nums)-1
	} else {
		left, right = 0, pivot-1
	} // else>

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		} // else>>
	} // for>
	return -1
}
