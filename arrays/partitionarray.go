package arrays

import "fmt"

func PartitionArr() {
	a := []int{1, 14, 5, 20, 4, 2, 15, 16, 17, 18, 54, 20, 87, 98, 3, 1, 32}
	low, high := 14, 20
	i, j := 0, len(a)-1

	for k := 0; k <= j; {
		if a[k] < low {
			a[k], a[i] = a[i], a[k]
			i++
			k++
		} else if a[k] > high {
			a[k], a[j] = a[j], a[k]
			j--
			// Don't increment k here — recheck the swapped-in value
		} else {
			k++
		}
	}
	fmt.Printf("%-60s %v\n", "Partitioned array", a)
}
