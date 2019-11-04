package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	// 0-1节点
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

func printList(pHead *ListNode) {
	for pHead != nil {
		if pHead.Next == nil {
			fmt.Printf("%v", pHead.Val)
		} else {
			fmt.Printf("%v->", pHead.Val)
		}
		pHead = pHead.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 0, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	pHead := swapPairs(head)

	printList(pHead)

}
