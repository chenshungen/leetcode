package main

import "fmt"

func testIsSymmetric101() {
	tree := NewBinaryTree(1)
	lvl1Left := NewTreeNode(2)
	lvl1Right := NewTreeNode(2)
	lvl2Left1, lvl2Right1, lvl2Left2, lvl2Right2 := NewTreeNode(3), NewTreeNode(4), NewTreeNode(4), NewTreeNode(3)

	lvl1Left.Left = lvl2Left1
	lvl1Left.Right = lvl2Right1

	lvl1Right.Left = lvl2Left2
	lvl1Right.Right = lvl2Right2

	tree.Root.Left = lvl1Left
	tree.Root.Right = lvl1Right

	fmt.Println(isSymmetric(tree.Root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
