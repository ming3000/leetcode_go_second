package leetcode_go_second

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil {
		slow = slow.Next
	}

	var pre *ListNode
	for slow != nil {
		nextBack := slow.Next
		slow.Next = pre
		pre = slow
		slow = nextBack
	}

	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	} // for>

	return true
}
