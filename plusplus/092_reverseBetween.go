package plusplus

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	sentry := &ListNode{Next: head}
	// m 的前一个节点
	preM := sentry
	for i := 1; i <= m-1; i++ {
		preM = preM.Next
	}

	// m 和 n 区间内反转
	var pre, cur *ListNode
	cur = preM.Next
	for i := m; i <= n; i++ {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	}

	// preM.Next 指向 m 节点，此时原 m 节点应当指向 n+1 节点
	preM.Next.Next = cur
	// 反转后的区间头节点
	preM.Next = pre
	return sentry.Next
}
