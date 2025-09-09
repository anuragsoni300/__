package tree

import "container/list"

func (bt *BTree) LevelOrderTraversal() {
	arr := [][]int{}
	bt.Root.lordertravel(0, &arr)
	PLog("Level Order Recursion", arr)

	arr = [][]int{}
	bt.Root.lordertravel(0, &arr)
	PLog("Level Order Loop", arr)
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

func (root *TNode) lordertravel2(arr [][]int) [][]int {
	if root == nil {
		return arr
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		level := []int{}

		for range levelSize {
			e := queue.Front()
			node := e.Value.(*TNode)
			level = append(level, node.Val)
			queue.Remove(e)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		arr = append(arr, level)
	}
	return arr
}
