package main

import "fmt"

func Insertion() {
	a := []int{1, 2, 3}
	index := 2
	a = append(a[:index], append([]int{4}, a[index:]...)...)
	fmt.Printf("%-30s %v\n", "Insertion", a)
}
