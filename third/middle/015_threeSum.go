package middle

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	ret := make([][]int, 0, len(nums))
	sort.Ints(nums)
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
				} // for>>>>
				for left < right && nums[right] == nums[right-1] {
					right--
				} // for>>>>
				left++
				right--
			} // else>>>
		} // for>>
	} // for>
	return ret
}
