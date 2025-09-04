package tree

func (bt *BTree) Search(val int) {
	bt.Root.search(val)
}

func (root *TNode) search(val int) {
	if root == nil {
		PLog("Not Found", val)
		return
	}
	if root.Val == val {
		PLog("Found", val)
		return
	}
	if val < root.Val {
		root.Left.search(val)
	} else {
		root.Right.search(val)
	}
}
