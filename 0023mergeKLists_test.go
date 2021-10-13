package leetcode

import "testing"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var head *ListNode

	head = mergeTwoLists(lists[0], mergeKLists(lists[1:]))

	return head
}

func TestMergeKLists(t *testing.T) {
	l1 := newLinkedList(1, 4, 5)
	l1.Print()
	l2 := newLinkedList(1, 3, 4)
	l2.Print()
	l3 := newLinkedList(2, 6)
	l3.Print()
	l4 := mergeKLists([]*ListNode{l1, l2, l3})
	l4.Print()
}
