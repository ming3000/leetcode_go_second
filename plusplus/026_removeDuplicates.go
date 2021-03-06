package plusplus

// 使用快慢指针来记录遍历的坐标。
// 开始时这两个指针都指向第一个数字
// 如果两个指针指的数字相同，则快指针向前走一步
// 如果不同，则两个指针都向前走一步
// 当快指针走完整个数组后，慢指针当前的坐标加 1 就是数组中不同数字的个数
// 实际上这就是双指针中的快慢指针。在这里快指针是读指针， 慢指针是写指针。
func removeDuplicates(nums []int) int {
	var left int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[left] {
			left++
			nums[left] = nums[i]
		} // if>>
	} // for>
	return left + 1
}
