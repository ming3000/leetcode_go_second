package leetcode_go_second

func partition(head *ListNode, x int) *ListNode {
	bigSentry, smallSentry := &ListNode{Next: head}, &ListNode{Next: head}
	bigTail, smallTail := bigSentry, smallSentry
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
	smallTail.Next = bigSentry.Next
	return smallSentry.Next
}
