package arrays

import "fmt"

func RotateImage() {
	matrix := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	n := len(matrix[0])

	for i := range n / 2 {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
	fmt.Printf("%-60s %v\n", "ZigZag", matrix)
}
