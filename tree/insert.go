package tree

func (bt *BTree) Insert(val int) *BTree {
	if bt.Root == nil {
		bt.Root = &TNode{Val: val}
	} else {
		bt.Root.insert(val)
	}
	return bt
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
