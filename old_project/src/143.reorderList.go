package main

import "fmt"

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	pLeft, pRight := split(head)
	head = pLeft
	pRight = reverseList(pRight)
	p1, p2 := pLeft, pRight

	for p1 != nil && p2 != nil {
		p1Temp := p1.Next
		p2Temp := p2.Next
		p2.Next = nil
		p1.Next = p2
		p2.Next = p1Temp
		p1 = p1Temp
		p2 = p2Temp
	}
	if p2 != nil {
		pCurr := head
		for pCurr.Next != nil {
			pCurr = pCurr.Next
		}
		pCurr.Next = p2
	}
}

func reverseList(head *ListNode) *ListNode {
	var pTop *ListNode
	pCurr := head
	for pCurr != nil {
		pTemp := pCurr.Next
		pCurr.Next = pTop
		pTop = pCurr
		pCurr = pTemp
	}

	return pTop
}

func printList(head *ListNode) {
	cur := head
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

func testReorderList143() {
	list := NewLinkedList()
	list.InsertToTail(1)
	list.InsertToTail(2)
	list.InsertToTail(3)
	list.InsertToTail(4)
	list.InsertToTail(5)
	// list.InsertToTail(6)
	// list.InsertToTail(7)
	// list.InsertToTail(8)
	// list.InsertToTail(9)
	// list.InsertToTail(10)
	list.Print()

	reorderList(list.head.Next)
	list.Print()
}
