package main

import "fmt"

func RotateByK() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate := 9
	for range rotate {
		temp := a[0]
		a = a[1:len(a)]
		a = append(a, temp)
	}
	fmt.Printf("%-30s %v\n", "Selection Sort", a)
}
