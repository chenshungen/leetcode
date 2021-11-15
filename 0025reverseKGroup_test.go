package leetcode

import "testing"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || (head != nil && head.Next == nil) {
		return head
	}
	var newHead *ListNode
	phead := head
	pprev := head
	piter := head.Next
	count := 1
	for piter != nil {
		if count%k == 0 {
			pprev.Next = nil
			if newHead == nil {
				newHead = reverseList(phead)
			} else {
				pt := newHead
				for pt != nil {
					if pt.Next == nil {
						break
					}
					pt = pt.Next
				}
				if pt != nil {
					pt.Next = reverseList(phead)
				}
			}
			pprev = piter
			piter = piter.Next
			phead = pprev
		} else {
			pprev = piter
			piter = piter.Next
		}
		count++
	}

	if phead != nil {
		if count%k == 0 {
			if newHead == nil {
				newHead = reverseList(phead)
			} else {
				pt := newHead
				for pt != nil {
					if pt.Next == nil {
						break
					}
					pt = pt.Next
				}
				pt.Next = reverseList(phead)
			}
		} else {
			if newHead == nil {
				newHead = phead
			} else {
				pt := newHead
				for pt != nil {
					if pt.Next == nil {
						break
					}
					pt = pt.Next
				}
				pt.Next = phead
			}
		}
	}
	return newHead
}

func TestReverseKGroup(t *testing.T) {
	l := newLinkedList(1)
	l = reverseKGroup(l, 1)
	l.Print()
}
