package arrays

import "fmt"

func RepeatedElement() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 1}
	slow, fast := 0, 0

	for {
		slow = a[slow]
		fast = a[a[fast]]
		if slow == fast {
			break
		}
	}
	fmt.Printf("%-60s %d\n", "Repeated element", a[slow])
}
