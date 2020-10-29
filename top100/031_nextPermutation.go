package middle

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	// 从后往前找到第一个降序的,相当于找到了我们的回溯点
	for i > -1 && nums[i+1] <= nums[i] {
		i--
	}

	if i > -1 {
		// 找到从右边起第一个大于nums[i]的，并将其和nums[i]进行交换
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		} // for>>
		nums[i], nums[j] = nums[j], nums[i]
	} // if>

	reverseArr(nums, i+1, len(nums)-1)
}

func reverseArr(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
