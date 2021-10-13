package leetcode

import "testing"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = AlignList(l1, l2)
	var l3 *ListNode
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		remainder := sum % 10
		carry = sum / 10
		l3 = TailInsert(l3, remainder)

		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		l3 = TailInsert(l3, carry)
	}

	return l3
}

func TestAddTownNum(t *testing.T) {
	// l1 := newRevLinkedList(2, 4, 3)
	l1 := newRevLinkedList(9, 9, 9, 9, 9, 9, 9)
	l1.Print()
	// l2 := newRevLinkedList(5, 6, 4)
	l2 := newRevLinkedList(9, 9, 9, 9)
	l2.Print()
	l3 := addTwoNumbers(l1, l2)
	l3.Print()
}
