package middle

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	big, small := &ListNode{Next: head}, &ListNode{Next: head}
	bigTail, smallTail := big, small

	for head != nil {
		if head.Val < x {
			smallTail.Next = head
			smallTail = smallTail.Next
		} else {
			bigTail.Next = head
			bigTail = bigTail.Next
		} // else>>
		head = head.Next
	} // for>
	bigTail.Next = nil
	smallTail.Next = big.Next
	return small.Next
}
