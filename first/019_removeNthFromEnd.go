package first

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{
		Val:  0,
		Next: head,
	}
	var pre, cur *ListNode
	cur = sentry
	count := 1
	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} // if>>
		head = head.Next
		count += 1
	} // for>
	pre.Next = cur.Next
	return sentry.Next
}
