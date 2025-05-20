package arrays

import "fmt"

func MSort() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}

	a = MergeSort(a)
	fmt.Printf("%-30s %v\n", "Merge Sort", a)
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	first := MergeSort(a[:len(a)/2])
	second := MergeSort(a[len(a)/2:])
	a = Merge(first, second)
	return a
}

func Merge(a, b []int) []int {
	final, i, j := []int{}, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for i < len(a) {
		final = append(final, a[i])
		i++
	}
	for j < len(b) {
		final = append(final, b[j])
		j++
	}
	return final
}
