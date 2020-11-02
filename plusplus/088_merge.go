package plusplus

// 这里要求原地修改，其实我们能只要从后往前比较，并从后往前插入即可。
// 我们需要三个指针：
// current 用于记录当前填补到那个位置了
// m 用于记录 nums1 数组处理到哪个元素了
// n 用于记录 nums2 数组处理到哪个元素了
func merge(nums1 []int, m int, nums2 []int, n int) {
	retIndex, i, j := m+n-1, m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[retIndex] = nums2[j]
			j--
		} else {
			nums1[retIndex] = nums1[i]
			i--
		} // else>>
		retIndex--
	} // for>
	for j >= 0 {
		nums1[retIndex] = nums2[j]
		retIndex--
		j--
	} // for>
}
