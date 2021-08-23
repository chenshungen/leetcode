package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	pCurr := head
	pivot := pCurr
	// 先找大于x的节点
	for pCurr != nil && pCurr.Val < x {
		pivot = pCurr
		pCurr = pCurr.Next
	}

	// 没有找到大于x的节点则找等于x的节点
	if pCurr == nil {
		return head
	}

	pPre := pCurr
	pCurr = pCurr.Next
	for pCurr != nil {
		// 小于x的节点 先删除
		if pCurr.Val < x {
			// 删除节点
			pDelete := pCurr
			pPre.Next = pCurr.Next
			pCurr = pPre.Next
			// 插入节点
			if pivot.Val >= x {
				pDelete.Next = pivot
				head = pDelete
				pivot = head
			} else {
				pDelete.Next = pivot.Next
				pivot.Next = pDelete
				pivot = pivot.Next
			}
		} else {
			pPre = pCurr
			pCurr = pCurr.Next
		}
	}

	return head
}

func reverseList(pList *ListNode) *ListNode {

	pCurr := pList
	var pTop *ListNode
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.Next
		pCurr.Next = pTop
		pTop = pCurr
		pCurr = pTemp
	}

	return pTop
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// m==n 并且只划分出一段
	if m == n && (m == 1 || m == listLen(head)) {
		return head
	}

	if n-(m-1) == listLen(head) {
		return reverseList(head)
	}

	var (
		part1   *ListNode
		part2   *ListNode
		part3   *ListNode
		last    *ListNode
		partTmp *ListNode
	)

	// 划分成两段
	if m == 1 {
		part1, part2 = splitInK(head, n)
		part1 = reverseList(part1)
		last = getTailNode(part1)
		last.Next = part2
		return part1
	}

	if n == listLen(head) {
		part1, part2 = splitInK(head, m-1)
		part2 = reverseList(part2)
		last = getTailNode(part1)
		last.Next = part2
		return part1
	}

	part1, partTmp = splitInK(head, m-1)
	if partTmp != nil {
		part2, part3 = splitInK(partTmp, n-(m-1))
	}
	part2 = reverseList(part2)
	last = getTailNode(part1)
	if part2 != nil {
		last.Next = part2
	}
	last = getTailNode(part1)
	last.Next = part3
	return part1
}

func listLen(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	pCurr := head
	for pCurr != nil {
		pCurr = pCurr.Next
		count++
	}
	return count
}

func getTailNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pCurr := head
	for pCurr.Next != nil {
		pCurr = pCurr.Next
	}
	return pCurr
}

func splitInK(head *ListNode, k int) (*ListNode, *ListNode) {
	count := 0
	pCurr := head
	for pCurr != nil {
		pre := pCurr
		pCurr = pCurr.Next
		count++
		if count == k {
			pre.Next = nil
			return head, pCurr
		}
	}
	return head, pCurr
}

// https://leetcode.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	lLen := listLen(head)
	if lLen < k {
		return head
	}

	first, second := splitInK(head, k)

	first = reverseList(first)

	lastNode := getTailNode(first)

	lastNode.Next = reverseKGroup(second, k)

	return first
}

//https://leetcode.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var leftK int
	llen := listLen(head)
	if k > llen {
		leftK = llen - (k % llen)
	} else {
		leftK = llen - k
	}
	first, second := splitInK(head, leftK)
	lastNode := getTailNode(second)
	if lastNode != nil {
		lastNode.Next = first
	} else {
		return first
	}
	return second
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

	head := &ListNode{Val: 3, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	// node2 := &ListNode{Val: 3, Next: nil}
	// node3 := &ListNode{Val: 2, Next: nil}
	// node4 := &ListNode{Val: 5, Next: nil}
	// node5 := &ListNode{Val: 2, Next: nil}

	head.Next = node1
	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	printList(head)
	//printList(reverseKGroup(head, 2))
	// printList(rotateRight(head, 5))
	//printList(reverseBetween(head, 1, 5))
	printList(partition(head, 2))

}
