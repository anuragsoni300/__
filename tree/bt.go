package tree

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

type BTree struct {
	Root *TNode
}

func (bt *BTree) InOrder() {
	bt.Root.inOrder()
}

func (bt *BTree) PreOrder() {
	bt.Root.preOrder()
}

func (bt *BTree) PostOrder() {
	bt.Root.postOrder()
}

func (bt *BTree) Insert(val int) *BTree {
	if bt.Root == nil {
		bt.Root = &TNode{Val: val}
	} else {
		bt.Root.insert(val)
	}
	return bt
}

func (bt *BTree) Print() {
	bt.InOrder()
	print("\n")
}

func (root *TNode) insert(val int) {
	if val <= root.Val {
		if root.Left == nil {
			root.Left = &TNode{Val: val}
		} else {
			root.Left.insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TNode{Val: val}
		} else {
			root.Right.insert(val)
		}
	}
}

func (root *TNode) inOrder() {
	if root == nil {
		return
	}
	root.Left.inOrder()
	print(root.Val, " ")
	root.Right.inOrder()
}

func (root *TNode) preOrder() {
	if root == nil {
		return
	}
	print(root.Val, " ")
	root.Left.preOrder()
	root.Right.preOrder()
	return
}

func (root *TNode) postOrder() {
	if root == nil {
		return
	}
	root.Left.postOrder()
	root.Right.postOrder()
	print(root.Val, " ")
	return
}
