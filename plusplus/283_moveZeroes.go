package plusplus

// 可以借助一个游标记录位置，然后遍历一次，将非 0 的原地修改，最后补 0 即可。
func moveZeroes(nums []int) {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		} // if>
	} // for>
	for i := index; i < len(nums); i++ {
		nums[index] = 0
		index++
	} // for>
}
