package main

import "fmt"

func Deletion() {
	a := []int{1, 2, 3, 4}
	index := 3
	a = append(a[:index], a[index+1:]...)
	fmt.Printf("%-30s %v\n", "Deletion", a)
}
