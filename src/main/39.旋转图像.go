package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i != j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Print(matrix)
}
