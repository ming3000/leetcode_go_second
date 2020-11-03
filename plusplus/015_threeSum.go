package plusplus

import "sort"

// 采用分治的思想找出三个数相加等于 0，
// 我们可以数组依次遍历，每一项 a[i]我们都认为它是最终能够用组成 0 中的一个数字，
// 那么我们的目标就是找到剩下的元素（除 a[i]）两个相加等于-a[i].
// 通过上面的思路，我们的问题转化为了给定一个数组，
// 找出其中两个相加等于给定值，我们成功将问题转换为了另外一道力扣的简单题目 1. 两数之和。
// 这个问题是比较简单的， 我们只需要对数组进行排序，然后双指针解决即可。
// 加上需要外层遍历依次数组，因此总的时间复杂度应该是 O(N^2)。
//
// 关键点解析
// 排序之后，用双指针
// 分治
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)

	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		} // if>>
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // if>>
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+nums[i] < 0 {
				left++
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				} // for>>>
				for left < right && nums[right] == nums[right-1] {
					right--
				} // for>>>
				left++
				right--
			} // else>>>
		} // for>>
	} // for>
	return ret
}
