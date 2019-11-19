package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]*ListNode)
	p1, p2 := headA, headB
	for p1 != nil {
		m[p1] = p1
		p1 = p1.Next
	}

	for p2 != nil {
		p, ok := m[p2]
		if ok {
			return p
		}
		p2 = p2.Next
	}

	return nil

}
