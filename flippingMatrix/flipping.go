package main

import (
	"fmt"
)

func flippingMatrix(matrix [][]int) int {
	n := len(matrix) / 2
	maxSum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			topLeft := matrix[i][j]
			topRight := matrix[i][2*n-j-1]
			bottomLeft := matrix[2*n-i-1][j]
			bottomRight := matrix[2*n-i-1][2*n-j-1]

			maxValue := max(topLeft, topRight, bottomLeft, bottomRight)
			maxSum += maxValue
		}
	}

	return maxSum
}

func max(a, b, c, d int) int {
	if a >= b && a >= c && a >= d {
		return a
	}
	if b >= a && b >= c && b >= d {
		return b
	}
	if c >= a && c >= b && c >= d {
		return c
	}
	return d
}

func main() {
	// Example usage:
	matrix := [][]int{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	result := flippingMatrix(matrix)
	fmt.Println("Maximum sum of submatrix:", result)
}
