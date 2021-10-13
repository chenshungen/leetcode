package leetcode

import "testing"

func swapPairs(head *ListNode) *ListNode {

	if head == nil || (head != nil && head.Next == nil) {
		return head
	}

	p1 := head
	p2 := head.Next
	pNewHead := p2.Next

	p2.Next = p1
	p1.Next = swapPairs(pNewHead)

	return p2
}

func TestSwapPairs(t *testing.T) {
	l := newLinkedList(1, 2, 3, 4)
	l.Print()

	l = swapPairs(l)
	l.Print()
}
