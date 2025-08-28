package linkedlist

func (List *Node) RemoveDuplicatesFromSorted() *Node {
	prev := &Node{}
	head := List
	for List != nil {
		if prev.Val == List.Val {
			prev.Next = List.Next
			List = List.Next
		} else {
			prev = List
			List = List.Next
		}
	}
	return head
}
