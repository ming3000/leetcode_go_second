package middle

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre, cur *ListNode
	var count = 1
	sentry := &ListNode{Next: head}
	cur = sentry
	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} // if>>
		count += 1
		head = head.Next
	} // for>
	pre.Next = cur.Next
	return sentry.Next
}
