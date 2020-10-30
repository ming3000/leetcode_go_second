package middle

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha, hb := headA, headB
	var la, lb int
	for ha != nil {
		ha = ha.Next
		la++
	}
	for hb != nil {
		hb = hb.Next
		lb++
	}

	var longHead, shortHead *ListNode
	var diff int
	if la > lb {
		diff = la - lb
		longHead = headA
		shortHead = headB
	} else {
		diff = lb - la
		longHead = headB
		shortHead = headA
	}

	for longHead != nil {
		if diff > 0 {
			longHead = longHead.Next
			diff--
			continue
		} // if>>
		if longHead != shortHead {
			longHead = longHead.Next
			shortHead = shortHead.Next
		} else {
			break
		} // else>>
	} // for>
	return longHead
}
