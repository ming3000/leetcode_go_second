package plusplus

// 这里我们可以增加一个 Map 记录已经遍历过的数字及其对应的索引值。
// 这样当遍历一个新数字的时候就去 Map 里查询 target 与该数的差值是否已经在前面的数字中出现过。
// 如果出现过，就找到了答案，就不必再往下继续执行了。
//
// 关键点
// 求和转换为求差
// 借助 Map 结构将数组中每个元素及其索引相互对应
// 以空间换时间，将查找时间从 O(N) 降低到 O(1)

func twoSum(nums []int, target int) []int {
	cm := make(map[int]int, len(nums))
	for i, v := range nums {
		diff := target - v
		if index, ok := cm[diff]; ok {
			return []int{i, index}
		}
		cm[v] = i
	} // for>
	return []int{-1, -1}
}
