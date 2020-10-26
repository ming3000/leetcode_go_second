package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	retIndex, i, j := m+n-1, m-1, n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[retIndex] = nums1[i]
			i--
		} else {
			nums1[retIndex] = nums2[j]
			j--
		} // else>>
		retIndex--
	} // for>
}
