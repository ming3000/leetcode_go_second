package leetcode_go_second

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var ret [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ret
		} // if>>
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // if>>
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left += 1
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right -= 1
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				} // for>>>>
				for left < right && nums[right] == nums[right-1] {
					right--
				} // for>>>>
				left += 1
				right -= 1
			} // else>>>
		} // for>>
	} // for>

	return ret
}
