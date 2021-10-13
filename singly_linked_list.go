package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func ListLength(list *ListNode) int {
	length := 0
	cur := list
	for nil != cur {
		length++
		cur = cur.Next
	}
	return length
}

func AlignList(list1, list2 *ListNode) (*ListNode, *ListNode) {
	l1, l2 := ListLength(list1), ListLength(list2)
	if l1 > l2 {
		for i := 0; i < (l1 - l2); i++ {
			list2 = TailInsert(list2, 0)
		}
	} else if l1 < l2 {
		for i := 0; i < (l2 - l1); i++ {
			list1 = TailInsert(list1, 0)
		}
	}

	return list1, list2
}

func HeadInsert(list *ListNode, val int) *ListNode {
	if list == nil {
		return newListNode(val)
	}
	phead := list
	tmpNode := newListNode(val)
	tmpNode.Next = phead
	list = tmpNode
	return list
}

func TailInsert(list *ListNode, val int) *ListNode {
	if list == nil {
		return newListNode(val)
	}

	tail := list
	for tail.Next != nil {
		tail = tail.Next
	}

	tmpNode := newListNode(val)
	tail.Next = tmpNode
	return list
}
