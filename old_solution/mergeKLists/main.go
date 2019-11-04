package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l2.Next, l1)
	}
	return head
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
	list1Node2 := &ListNode{Val: 5, Next: nil}
	list1Node1 := &ListNode{Val: 4, Next: list1Node2}
	list1 := &ListNode{Val: 1, Next: list1Node1}

	list2Node2 := &ListNode{Val: 4, Next: nil}
	list2Node1 := &ListNode{Val: 3, Next: list2Node2}
	list2 := &ListNode{Val: 1, Next: list2Node1}

	list3Node1 := &ListNode{Val: 6, Next: nil}
	list3 := &ListNode{Val: 2, Next: list3Node1}

	lists := make([]*ListNode, 0, 4)
	lists = append(lists, list1, list2, list3)

	sorted := mergeKLists(lists)

	printList(sorted)
}
