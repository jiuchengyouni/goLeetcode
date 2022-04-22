package main

import (
	"fmt"
)

func setMatrix(n int, m int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	return grid
}
func transpose(grid [][]int, n int, m int) [][]int {
	grid_r := make([][]int, m)
	for i := 0; i < m; i++ {
		grid_r[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid_r[j][i] = grid[i][j]
		}
	}
	return grid_r
}
func main() {
	var (
		n int
		m int
		k int
	)
	fmt.Scan(&n, &m, &k)
	//
	grid := setMatrix(n, m)
	grid_r := transpose(grid, n, m)
	//
	grid2 := setMatrix(m, k)
	grid_r2 := transpose(grid2, m, k)
	grid3 := make([][]int, m)
	for i := 0; i < m; i++ {
		grid3[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for y := 0; y < k; y++ {
				grid3[i][j] += grid_r[i][y] * grid_r2[y][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			print(grid3[i][j])
			print(" ")
		}
		fmt.Println()
	}
}
