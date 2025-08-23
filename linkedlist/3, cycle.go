package linkedlist

import "fmt"

func (list *Linklist) _middle() {
	slow, fast := list.head, list.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Printf("%-60s %v\n", "Middle element", slow.Val)
}

func (list *Linklist) _detectcycle() {
	slow, fast := list.head, list.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fmt.Println("Cycle prsent")
			break
		}
	}

	if slow == fast {
		slow = list.head
		var prev *Node
		for slow != fast {
			prev = fast
			slow = slow.Next
			fast = fast.Next
		}
		prev.Next = nil
	}
}
