package leetcode_go_second

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	midPre, mid := findMidPre(head)
	if midPre == nil {
		midPre.Next = nil
	}

	root := &TreeNode{Val: mid.Val}
	if head == mid {
		return root
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func findMidPre(head *ListNode) (*ListNode, *ListNode) {
	var pre, slow, fast *ListNode
	slow, fast = head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre, slow
}
