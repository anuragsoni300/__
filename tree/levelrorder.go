package tree

func (bt *BTree) LevelOrderTraversal() {
	arr := [][]int{}
	bt.Root.lordertravel(0, &arr)
	PLog("Level Order", arr)
}

func (root *TNode) lordertravel(level int, arr *[][]int) {
	if root == nil {
		return
	}
	if len(*arr) == level {
		*arr = append(*arr, []int{})
	}
	(*arr)[level] = append((*arr)[level], root.Val)
	root.Left.lordertravel(level+1, arr)
	root.Right.lordertravel(level+1, arr)
}
