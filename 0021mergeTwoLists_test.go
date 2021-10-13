package leetcode

import "testing"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var head, iter *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp := l1
			l1 = l1.Next
			tmp.Next = nil
			if head == nil {
				head = tmp
				iter = head
			} else if iter != nil {
				iter.Next = tmp
				iter = iter.Next
			}
		} else {
			tmp := l2
			l2 = l2.Next
			tmp.Next = nil
			if head == nil {
				head = tmp
				iter = head
			} else if iter != nil {
				iter.Next = tmp
				iter = iter.Next
			}
		}
	}

	if l1 != nil {
		iter.Next = l1
	} else if l2 != nil {
		iter.Next = l2
	}

	return head
}

func TestMergeTwoLists(t *testing.T) {
	l1 := newLinkedList(1, 2, 4)
	l1.Print()
	l2 := newLinkedList(1, 3, 4)
	l2.Print()
	l3 := mergeTwoLists2(l1, l2)
	l3.Print()
}
