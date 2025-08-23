package linkedlist

func (list *Linklist) _reverse() {
	curr := list.head
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func (list *Node) _reverseRec() *Node {
	if list == nil || list.Next == nil {
		return list
	}
	rest := list.Next._reverseRec()
	list.Next.Next = list
	list.Next = nil
	return rest
}
