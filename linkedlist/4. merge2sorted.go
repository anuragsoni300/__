package linkedlist

func (list1 *Node) _mergeSortedRec(list2 *Node) *Node {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = list1.Next._mergeSortedRec(list2)
		return list1
	} else {
		list2.Next = list2.Next._mergeSortedRec(list1)
		return list2
	}
}

func (list1 *Node) _mergeSortedItr(list2 *Node) *Node {
	dummy := &Node{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return dummy.Next
}
