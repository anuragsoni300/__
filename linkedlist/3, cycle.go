package linkedlist

func (list *Linklist) Middle() {
	slow, fast := list.Head, list.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	PLog("Middle element", slow.Val)
}

func (list *Linklist) Detectcycle() {
	slow, fast := list.Head, list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			PLog("Cycle prsent and removing it", "")
			break
		}
	}

	if slow == fast {
		slow = list.Head
		var prev *Node
		for slow != fast {
			prev = fast
			slow = slow.Next
			fast = fast.Next
		}
		prev.Next = nil
	}
}
