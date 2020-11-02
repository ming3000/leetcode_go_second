package plusplus

// 使用栈，遍历输入字符串
// 如果当前字符为左半边括号时，则将其压入栈中
// 如果遇到右半边括号时，分类讨论：
// 1）如栈不为空且为对应的左半边括号，则取出栈顶元素，继续循环
// 2）若此时栈为空，则直接返回 false
// 3）若不为对应的左半边括号，反之返回 false

func isValid(s string) bool {
	cs := make([]byte, 0, len(s))
	cm := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			cs = append(cs, s[i])
		} else {
			if len(cs) == 0 {
				return false
			} // if>>>
			top := cs[len(cs)-1]
			cs = cs[:len(cs)-1]
			if cm[top] != s[i] {
				return false
			} //if>>>
		} // else>>
	} // for>
	return len(cs) == 0
}
