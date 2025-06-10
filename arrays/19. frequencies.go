package arrays

import "fmt"

func Frequencies() {
	a := []int{11, 2, 3, 3, 3, 44, 5, 5, 5, 5, 5, 5, 6, 6}
	mp := make(map[int]int)

	for _, v := range a {
		mp[v]++
	}

	fmt.Printf("%-60s %v\n", "Frequencies", mp)
	fmt.Printf("%-60s %v\n", "Count", len(mp))
}
