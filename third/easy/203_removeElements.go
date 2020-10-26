package easy

func removeElements(head *ListNode, val int) *ListNode {
	sentry := &ListNode{Next: head}
	pre, cur := sentry, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		} // else>>
		cur = cur.Next
	} // for>
	return sentry.Next
}
