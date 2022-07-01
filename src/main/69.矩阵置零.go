package main

import "fmt"

func setZeroes(matrix [][]int) {
	ans := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				ans1 := []int{i, j}
				ans = append(ans, ans1)
			}
		}
	}
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[ans[i][0]][j] = 0
		}
	}
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][ans[i][1]] = 0
		}
	}
}
func main() {
	matrix := [][]int{{0, 1, 2, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
