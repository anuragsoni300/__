package arrays

import "fmt"

func RemoveDuplicates() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	j := 1

	for i := 1; i < len(a) && j < len(a); {
		if a[i] == a[j-1] {
			i++
		} else {
			a[i], a[j] = a[j], a[i]
			i++
			j++
		}
	}
	fmt.Printf("%-60s %v\n", "Remove duplicates", a)
}
