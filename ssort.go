package main

import "fmt"

func SSort() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}

	for i := range a {
		j := i + 1
		min, minIndex := a[i], i
		for j < len(a) {
			if a[j] < min {
				min = a[j]
				minIndex = j
			}
			j++
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	fmt.Printf("%-30s %v\n", "Selection Sort", a)
}
