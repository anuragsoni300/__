package linkedlist

func (end *Node) _pallindrome(start **Node) bool {
	if end == nil {
		return true
	}

	if !end.Next._pallindrome(start) || end.Val != (*start).Val {
		return false
	}
	*start = (*start).Next
	return true
}

func (List *Node) IsPallindromeRec() {
	start := &List
	if List._pallindrome(start) {
		PLog("Pallindrome", "Yes")
	} else {
		PLog("Pallindrome", "No")
	}
}

func (List *Node) IsPallindromeItr() {
	slow, fast := List, List

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	slow = prev
	curr = List
	isPallindrome := true
	for slow != nil && curr != nil {
		if slow.Val != curr.Val {
			isPallindrome = false
		}
		slow = slow.Next
		curr = curr.Next
	}

	if isPallindrome {
		PLog("Pallindrome", "Yes")
	} else {
		PLog("Pallindrome", "No")
	}
}
