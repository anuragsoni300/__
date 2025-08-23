package linkedlist

func (list *Linklist) _delete(val int) {
	curr := list.head
	prev := curr
	for curr != nil && curr.Val != val {
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next
}
