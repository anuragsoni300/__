package linkedlist

func (List *Node) RemoveDuplicatesFromSorted() *Node {
	head := List
	for List != nil && List.Next != nil {
		if List.Val == List.Next.Val {
			List.Next = List.Next.Next
		} else {
			List = List.Next
		}
	}
	return head
}
