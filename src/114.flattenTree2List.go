package main

func testFlatten114() {
	tree := NewBinaryTree(1)
	lvl1Left := NewTreeNode(2)
	lvl1Right := NewTreeNode(5)
	lvl2Left1, lvl2Right1, lvl2Left2, lvl2Right2 := NewTreeNode(3), NewTreeNode(4), NewTreeNode(6), NewTreeNode(7)

	lvl1Left.Left = lvl2Left1
	lvl1Left.Right = lvl2Right1

	lvl1Right.Left = lvl2Left2
	lvl1Right.Right = lvl2Right2

	tree.Root.Left = lvl1Left
	tree.Root.Right = lvl1Right

	flatten(tree.Root)

	tree.InOrderTraverse()
}

func flatten(root *TreeNode) {
	recur(root)
}

// recur 会 flatten root，并返回 flatten 后的 leaf 节点
func recur(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right
	root.Right = root.Left
	// 不要忘记封闭 left 节点
	root.Left = nil

	return res
}
