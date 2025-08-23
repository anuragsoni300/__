package linkedlist

func (list *Linklist) Delete(val int) {
	curr := list.Head
	prev := curr
	for curr != nil && curr.Val != val {
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next
}
