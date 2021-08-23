package main

import "fmt"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	if ld-rd > 1 || ld-rd < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func testIsBalanced110() {
	tree := NewBinaryTree(1)
	tree.Root.Left = NewTreeNode(2)
	tree.Root.Right = NewTreeNode(2)

	tree.Root.Left.Left = NewTreeNode(3)
	tree.Root.Left.Right = NewTreeNode(3)

	tree.Root.Left.Left.Left = NewTreeNode(4)

	fmt.Println(isBalanced(tree.Root))
}
