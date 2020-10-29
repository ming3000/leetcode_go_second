package first

func rotate(nums []int, k int) {
	swapArr(nums)
	swapArr(nums[:k%len(nums)])
	swapArr(nums[k%len(nums):])
}

func swapArr(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
