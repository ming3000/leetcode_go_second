package first

func deleteDuplicates(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	pre, cur := sentry, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		} // for>>
		if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = cur
		} // else>>
		cur = cur.Next
	} // for>
	return sentry.Next
}
