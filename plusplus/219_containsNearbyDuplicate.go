package plusplus

// 由于题目没有对空间复杂度有求，用一个hashmap 存储已经访问过的数字即可,
// 每次访问都会看hashmap中是否有这个元素，有的话拿出索引进行比对，
// 是否满足条件（相隔不大于k），如果满足返回true即可。
func containsNearbyDuplicate(nums []int, k int) bool {
	cm := make(map[int]bool, len(nums))
	for i, v := range nums {
		if _, ok := cm[v]; ok {
			return true
		} // if>>
		cm[v] = true
		if len(cm) > k {
			delete(cm, nums[i-k])
		}
	} // for>
	return false
}
