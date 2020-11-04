package plusplus

//初始化快慢指针 slow ， fast ，全部指向索引为 0 的元素。
//fast 每次移动一格
//慢指针选择性移动，即只有写入数据之后才移动。是否写入数据取决于 slow - 2 对应的数字和 fast 对应的数字是否一致。
//如果一致，我们不应该写。 否则我们就得到了三个相同的数字，不符合题意
//如果不一致，我们需要将 fast 指针的数据写入到 slow 指针。
//重复这个过程，直到 fast 走到头，说明我们已无数字可写。
func removeDuplicates80(nums []int) int {
	const k = 2
	var slow = 0
	for fast := 0; fast < len(nums); fast++ {
		if slow < k || nums[fast] != nums[slow-k] {
			nums[slow] = nums[fast]
			slow++
		} // if>
	} // for>
	return slow
}
