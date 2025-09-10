package tree

func (bt *BTree) Height() {
	height := bt.Root.maxHeight()
	PLog("Height", height)
}

func (root *TNode) maxHeight() int {
	if root == nil {
		return 0
	}

	left := root.Left.maxHeight() + 1
	right := root.Right.maxHeight() + 1

	if left > right {
		return left
	} else {
		return right
	}
}
