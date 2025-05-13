package main

import "fmt"

func BSearch() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	search := 5
	i, j := 0, len(a)-1

	for i <= j {
		mid := i + (j-i)/2
		if a[mid] == search {
			fmt.Printf("%-30s %d\n", "Found at", mid)
			return
		} else if a[mid] > search {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	fmt.Println("Not Exist")
}
