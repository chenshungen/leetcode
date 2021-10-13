package leetcode

import "testing"

func reverseList(head *ListNode) *ListNode {

	if head == nil || (head != nil && head.Next == nil) {
		return head
	}

	p1 := head
	p2 := head.Next
	head.Next = nil
	var p3 *ListNode

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}

func reverseList2(head *ListNode) *ListNode {

	if head == nil || (head != nil && head.Next == nil) {
		return head
	}

	var prv, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prv
		prv = head
		head = next
	}

	return prv
}

func reverseList3(head *ListNode) *ListNode {
	if head == nil || (head != nil && head.Next == nil) {
		return head
	}

	newHead := reverseList3(head.Next)

	head.Next.Next = head

	head.Next = nil

	return newHead
}

func TestReverseList(t *testing.T) {
	l := newLinkedList(1, 2, 3, 4, 5, 6)
	l = reverseList2(l)
	l.Print()
}
