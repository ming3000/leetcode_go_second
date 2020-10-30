package middle

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	r := sortList(mid)
	l := sortList(head)
	sentry := &ListNode{}
	return mergeList(sentry, l, r)
}

func mergeList(sentry, left, right *ListNode) *ListNode {
	tail := sentry
	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		} // else>>
		tail = tail.Next
	} // for>
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}
	return sentry.Next
}
