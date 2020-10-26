package middle

func swapPairs(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	cur := sentry
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second
		cur = first
	}

	return sentry.Next
}
