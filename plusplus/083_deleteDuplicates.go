package plusplus

func deleteDuplicates(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	pre, cur := sentry, head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		} // else>>
		cur = cur.Next
	} // for>
	return sentry.Next
}
