package arrays

import "fmt"

func hoare(a []int, low, high int) int {
	pivot, i, j := a[low], low, high
	for {
		for ; a[i] < pivot; i++ {
		}
		for ; a[j] > pivot; j-- {
		}
		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}

func quickNaive(a []int) []int {
	if len(a) < 2 {
		return a
	}
	less, greater, pivot := []int{}, []int{}, a[0]
	for _, v := range a[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(quickNaive(less), pivot), quickNaive(greater)...)
}

func lomuto(a []int, low, high int) int {
	i, pivot := low, a[high-1]
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high-1] = a[high-1], a[i]
	return i
}

func quickSort(a []int, low, high int) {
	if low < high {
		p := hoare(a, low, high)
		quickSort(a, low, p)
		quickSort(a, p+1, high)
	}
}

func quickSort_lomuto(a []int, low, high int) {
	if low < high {
		p := lomuto(a, low, high)
		quickSort_lomuto(a, low, p)
		quickSort_lomuto(a, p+1, high)
	}
}

func QSort() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}
	low, high := 0, len(a)-1
	quickSort(a, low, high)
	fmt.Printf("%-60s %v\n", "Quick Hoare", a)

	a = []int{7, 8, 9, 4, 5, 3, 1, 2, 6}
	a = quickNaive(a)
	fmt.Printf("%-60s %v\n", "Quick Naive", a)

	a = []int{7, 8, 9, 4, 5, 3, 1, 2, 6}
	low, high = 0, len(a)
	quickSort_lomuto(a, low, high)
	fmt.Printf("%-60s %v\n", "Quick Loumto", a)
}
