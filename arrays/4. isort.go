package arrays

import "fmt"

func ISort() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}

	for i := range a {
		j := i + 1
		for j > 0 && j < len(a) && a[j] < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}
	fmt.Printf("%-60s %v\n", "Insertion Sort", a)
}
