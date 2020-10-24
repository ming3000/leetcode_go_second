package leetcode_go_second

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		if nums[low] <= nums[high] {
			break
		}
		mid = (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	} // for>
	return nums[low]
}
