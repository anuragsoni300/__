package linkedlist

func IntersectionPoint1(list1, list2 *Node) {
	mem := make(map[Node]bool)
	curr := list1

	for curr != nil {
		mem[*curr] = true
		curr = curr.Next
	}

	curr = list2
	for curr != nil {
		if mem[*curr] {
			PLog("Intersection point using hashMap", curr.Val)
			break
		}
		curr = curr.Next
	}
}

func IntersectionPoint2(list1, list2 *Node) {
	len1, len2 := 0, 0
	curr1 := list1
	for curr1 != nil {
		len1++
		curr1 = curr1.Next
	}
	curr2 := list2
	for curr2 != nil {
		len2++
		curr2 = curr2.Next
	}

	curr1 = list1
	curr2 = list2

	if len1 < len2 {
		final := len2 - len1
		for final > 0 {
			final--
			curr2 = curr2.Next
		}
	} else {
		final := len1 - len2
		for final > 0 {
			final--
			curr1 = curr1.Next
		}
	}

	for curr1 != nil && curr2 != nil {
		if curr1 == curr2 {
			PLog("Intersection point using Node count", curr1.Val)
			break
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
}
