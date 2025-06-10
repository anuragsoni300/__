package arrays

import "fmt"

func Deletion() {
	a := []int{1, 2, 3, 4}
	index := 3
	a = append(a[:index], a[index+1:]...)
	fmt.Printf("%-60s %v\n", "Deletion", a)
}
