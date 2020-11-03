package plusplus

// 设置一个 dummy 节点简化操作，dummy next 指向 head。
// 初始化 first 为第一个节点
// 初始化 second 为第二个节点
// 初始化 current 为 dummy
// first.next = second.next
// second.next = first
// current.next = second
// current 移动两格
func swapPairs(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	cur := sentry
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		first, second := cur.Next, cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second
		cur = first
	}
	return sentry.Next
}
