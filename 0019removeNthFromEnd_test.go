package leetcode

import "testing"

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

	llen := ListLength(head)

	if n > llen {
		return nil
	}

	prv := head
	cur := head
	for i := 0; i < (llen - n); i++ {
		if cur != prv {
			prv = prv.Next
		}
		cur = cur.Next
	}

	if prv == cur {
		if cur.Next == nil {
			return nil
		} else {
			head = cur.Next
		}
	}

	if cur.Next == nil {
		prv.Next = nil
	} else {
		pnext := cur.Next
		prv.Next = pnext
	}

	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	l := newLinkedList(1, 2)
	l.Print()
	l = removeNthFromEnd(l, 2)
	l.Print()
}
