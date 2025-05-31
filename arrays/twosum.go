package arrays

import "fmt"

func TwoSum() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	mp := make(map[int]int, len(a))
	target := 10
	ans := make([]int, 2)
	for i, v := range a {
		if j, ok := mp[target-v]; ok {
			ans = []int{j, i}
			break
		} else {
			mp[v] = i
		}
	}
	fmt.Printf("%-60s %v\n", "Two sum index", ans)
}
