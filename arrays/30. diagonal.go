package arrays

import "fmt"

func DiagonalTraversal() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var ans []int
	row, column := 0, 0
	m, n := len(a), len(a[0])
	up := true

	for len(ans) < m*n {
		ans = append(ans, a[row][column])
		if up {
			if column == n-1 {
				row++
				up = false
			} else if row == 0 {
				column++
				up = false
			} else {
				row--
				column++
			}
		} else {
			if row == m-1 {
				column++
				up = true
			} else if column == 0 {
				row++
				up = true
			} else {
				row++
				column--
			}
		}
	}
	fmt.Printf("%-60s %v\n", "ZigZag", ans)
}
