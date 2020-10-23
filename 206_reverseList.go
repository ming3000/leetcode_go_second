package leetcode_go_second

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		nextBack := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextBack
	}
	return pre
}
