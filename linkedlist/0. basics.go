package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type Linklist struct {
	head *Node
}

func (list *Linklist) _insert(val int) {
	newNode := &Node{Val: val}
	if list.head == nil {
		list.head = newNode
	} else {
		curr := list.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
}

func (list *Linklist) _insertAfter(v1, v2 int) {
	newNode := &Node{Val: v2}

	curr := list.head
	for curr != nil {
		if curr.Val == v1 {
			break
		}
		curr = curr.Next
	}
	temp := curr.Next
	curr.Next = newNode
	newNode.Next = temp
}

func (list *Linklist) _delete(val int) {
	curr := list.head
	prev := curr
	for curr != nil && curr.Val != val {
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next
}

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

func (list *Linklist) _length() {
	count := 0
	curr := list.head
	for curr != nil {
		count++
		curr = curr.Next
	}
	fmt.Printf("%-60s %v\n", "Length of thr linked list", count)
}

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

func (list *Linklist) _print() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%v->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func Linkedlist() {
	myList, list2 := Linklist{}, Linklist{}
	myList._insert(1)
	// myList._insert(2)
	myList._insert(3)
	// myList._insert(4)
	myList._insert(5)
	// myList._insert(6)
	myList._insert(7)
	// myList._insert(8)
	myList._insert(9)
	// myList._insert(10)

	list2._insert(2)
	list2._insert(4)
	list2._insert(6)
	list2._insert(8)

	list2._insert(10)
	// myList._insertAfter(2, 5)
	// myList._delete(3)
	// myList._print()
	// myList.head = myList.head._reverseRec()
	// myList._reverse()
	// myList._print()
	// myList._length()
	// myList._middle()

	// curr := myList.head
	// for curr.Next != nil {
	// 	curr = curr.Next
	// }
	// curr.Next = myList.head.Next
	// myList._detectcycle()
	// myList._print()
	// myList.head = myList.head._mergeSortedRec(list2.head)
	myList.head = myList.head._mergeSortedItr(list2.head)
	myList._print()
}
