package middle

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre, cur *ListNode
	sentry := &ListNode{Next: head}
	count := 1
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
