package arrays

import "fmt"

func BSort() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}
	i, j := 0, len(a)

	for j > 0 {
		i = 1
		for i < j {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
			}
			i++
		}
		j--
	}
	fmt.Printf("%-30s %v\n", "Bubble Sort", a)
}
