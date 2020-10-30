package middle

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	} // for>
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}

	var pre *ListNode
	for slow != nil {
		nextBack := slow.Next
		slow.Next = pre
		pre = slow
		slow = nextBack
	}

	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		} // if>>
		head = head.Next
		pre = pre.Next
	} // for>
	return true
}
