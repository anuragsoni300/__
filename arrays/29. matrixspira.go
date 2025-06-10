package arrays

import "fmt"

func MatrixSpiral() {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	m, n := len(a), len(a[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	ans := []int{}

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans = append(ans, a[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			ans = append(ans, a[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, a[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, a[i][left])
			}
			left++
		}
	}
	fmt.Printf("%-60s %v\n", "Two sum index", ans)
}
