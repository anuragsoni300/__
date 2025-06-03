package arrays

import "fmt"

func ColorSort() {
	a := []int{2, 0, 2, 1, 1, 0}
	i, j, m := 0, len(a)-1, 0
	for m <= j {
		if a[m] == 0 {
			a[i], a[m] = a[m], a[i]
			i++
			m++
		} else if a[m] == 1 {
			m++
		} else {
			a[j], a[m] = a[m], a[j]
			j--
		}
	}
	fmt.Printf("%-60s %v\n", "Colors partition", a)
}
