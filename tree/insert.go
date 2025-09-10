package tree

import "container/list"

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

func (bt *BTree) insertlevelorder(val int) {
	if bt.Root == nil {
		bt.Root = &TNode{Val: val}
	}
	queue := list.New()
	queue.PushBack(bt.Root)
	for queue.Len() > 0 {
		levelSize := queue.Len()

		for range levelSize {
			node := queue.Front()
			queue.Remove(node)
			root := node.Value.(*TNode)
			if root.Left != nil {
				queue.PushBack(root.Val)
			} else {
				root.Left = &TNode{Val: val}
			}

			if root.Right != nil {
				queue.PushBack(root.Val)
			} else {
				root.Right = &TNode{Val: val}
			}
		}
	}
}
