package plusplus

// 设置虚拟节点 dummyHead 指向 head（简化判断，使得头结点不需要特殊判断）
// 设定双指针 p 和 q，初始都指向虚拟节点 dummyHead
// 移动 q，直到 p 与 q 之间相隔的元素个数为 n
// 同时移动 p 与 q，直到 q 指向的为 NULL
// 将 p 的下一个节点指向下下个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{Next: head}
	var count = 1
	var pre, cur *ListNode
	cur = sentry
	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} // if>>
		count++
		head = head.Next
	} // for>
	pre.Next = cur.Next
	return sentry.Next
}
