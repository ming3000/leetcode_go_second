package leetcode_go_second

func deleteDuplicates083(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		} // else>>
	} // for>

	return sentry.Next
}
