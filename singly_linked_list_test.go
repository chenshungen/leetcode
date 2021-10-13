package leetcode

import (
	"fmt"
	"testing"
)

func newRevLinkedList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	phead := newListNode(vals[0])
	for i := 1; i < len(vals); i++ {
		phead = HeadInsert(phead, vals[i])
	}

	return phead
}

func newLinkedList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	phead := newListNode(vals[0])
	for i := 1; i < len(vals); i++ {
		phead = TailInsert(phead, vals[i])
	}

	return phead
}

//Print 打印链表
func (list *ListNode) Print() {
	cur := list
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func TestSinglyLinkedList(t *testing.T) {
	// l1 := newRevLinkedList(1, 2, 3)
	// l1.Print()
	// l2 := newRevLinkedList(1, 2, 3, 4, 5, 6)
	// l2.Print()
	// l1, l2 = AlignList(l1, l2)
	// l1.Print()
	// l2.Print()

	// l3 := newLinkedList(1, 2, 3)
	// l3.Print()
	// l4 := newLinkedList(1, 2, 3, 4, 5, 6)
	// l4.Print()
	// l3, l4 = AlignList(l3, l4)
	// l3.Print()
	// l4.Print()

	l := newLinkedList(1, 2, 3, 4, 5, 6)
	l = reverseList2(l)
	l.Print()
}
