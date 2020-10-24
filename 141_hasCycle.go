package leetcode_go_second

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		} // if>>
		slow = slow.Next
		fast = fast.Next.Next
	} // for>
	return false
}
