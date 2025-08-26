package linkedlist

func (List *Node) NthNodeFromLast(n int) {
	main_ptr, ref_ptr := List, List

	for n > 0 && ref_ptr != nil {
		ref_ptr = ref_ptr.Next
		n--
	}

	for ref_ptr != nil {
		main_ptr = main_ptr.Next
		ref_ptr = ref_ptr.Next
	}

	if main_ptr != nil && n != 1 {
		PLog("Nth node from last", main_ptr.Val)
	} else {
		PLog("Wrong Index", "")
	}
}
