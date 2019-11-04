package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1 := head
	p2 := head.Next
	if p2 == nil {
		return false
	}

	// p1向后走一步 p2向后走两步 如果存在环 则p2必定能追上p1(操场跑圈的原理)
	for {
		if p1 == nil || p2 == nil {
			return false
		}

		if p1 == p2 {
			return true
		}

		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return false
		}
		p2 = p2.Next
	}
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	nodeMap := make(map[*ListNode]*ListNode)
	p := head
	for p != nil {
		node, ok := nodeMap[p]
		if !ok {
			nodeMap[p] = p
		} else {
			return node
		}
		p = p.Next
	}

	return nil
}

func main() {
	head := &ListNode{Val: 0, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}

	// head.Next = node1
	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(head))

	node := detectCycle(head)
	fmt.Println(node.Val)
}
