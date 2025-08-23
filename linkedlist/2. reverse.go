package linkedlist

func (list *Linklist) Reverse() {
	curr := list.Head
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.Head = prev
}

func (list *Node) ReverseRec() *Node {
	if list == nil || list.Next == nil {
		return list
	}
	rest := list.Next.ReverseRec()
	list.Next.Next = list
	list.Next = nil
	return rest
}
