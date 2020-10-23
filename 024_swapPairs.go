package leetcode_go_second

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentry := &ListNode{Next: head}
	pre := sentry
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		first.Next = second.Next
		second.Next = first
		pre.Next = second
		pre = first
	} // for>
	return sentry.Next
}
