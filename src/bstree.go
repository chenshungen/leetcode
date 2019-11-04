package main

// BST 二叉查找树
type BST struct {
	*BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

// NewBST 创建一棵二叉查找树
func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

// Find  查找
func (bst *BST) Find(v interface{}) *TreeNode {
	p := bst.Root
	for p != nil {
		compareResult := bst.compareFunc(v, p.Val)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	return nil
}

// Insert 插入
func (bst *BST) Insert(v interface{}) bool {
	p := bst.Root
	for p != nil {
		compareResult := bst.compareFunc(v, p.Val)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if p.Right == nil {
				p.Right = NewTreeNode(v)
				break
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = NewTreeNode(v)
				break
			}
			p = p.Left
		}
	}
	return true
}

// Delete 删除
func (bst *BST) Delete(v interface{}) bool {
	var pp *TreeNode
	p := bst.Root
	deleteLeft := false
	for nil != p {
		compareResult := bst.compareFunc(v, p.Val)
		if compareResult > 0 {
			pp = p
			p = p.Right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.Left
			deleteLeft = true
		} else {
			break
		}
	}

	if nil == p { //需要删除的节点不存在
		return false
	} else if nil == p.Left && nil == p.Right { //删除的是一个叶子节点
		if nil != pp {
			if deleteLeft {
				pp.Left = nil
			} else {
				pp.Right = nil
			}
		} else { //根节点
			bst.Root = nil
		}
	} else if nil != p.Right { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p
		q := p.Right //向右走一步
		fromRight := true
		for nil != q.Left { //向左走到底
			pq = q
			q = q.Left
			fromRight = false
		}
		if fromRight {
			pq.Right = nil
		} else {
			pq.Left = nil
		}
		q.Left = p.Left
		q.Right = p.Right
		if nil == pp { //根节点被删除
			bst.Root = q
		} else {
			if deleteLeft {
				pq.Left = q
			} else {
				pq.Right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.Left = p.Left
			} else {
				pp.Right = p.Left
			}
		} else {
			if deleteLeft {
				bst.Root = p.Left
			} else {
				bst.Root = p.Left
			}
		}
	}

	return true

}
