package plusplus

// 用一个 hashmap 来建立字符和其出现位置之间的映射。
// 同时维护一个滑动窗口，窗口内的都是没有重复的字符，去尽可能的扩大窗口的大小，窗口不停的向右滑动。
// 如果当前遍历到的字符从未出现过，那么直接扩大右边界；
// 如果当前遍历到的字符出现过，则缩小窗口（左边索引向右移动），然后继续观察当前遍历到的字符；
// 重复（1）（2），直到窗口内无重复元素；
// 维护一个全局最大窗口 res，每次用出现过的窗口大小来更新结果 res，最后返回 res 获取结果;
// 最后返回 res 即可;
func lengthOfLongestSubstring(s string) int {
	cm := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		cm[s[i]] = 0
	}

	var left int
	var ret int
	for i := 0; i < len(s); i++ {
		for cm[s[i]] != 0 {
			cm[s[left]]--
			left++
		} // for>>
		cm[s[i]]++
		ret = maxInt(ret, i-left+1)
	} // for>
	return ret
}
