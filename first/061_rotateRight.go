package first

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	pre := head
	length := 1
	for pre.Next != nil {
		pre = pre.Next
		length += 1
	}
	pre.Next = head

	cur := head
	for i := 0; i < length-k%length-1; i++ {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil
	return newHead
}
