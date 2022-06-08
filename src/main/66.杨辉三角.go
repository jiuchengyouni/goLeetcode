package main

import "fmt"

func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		arr[i-1] = make([]int, i)
	}
	if numRows == 0 {
		return nil
	}
	for i := 0; i < numRows; i++ {
		arr[i][0] = 1
		arr[i][i] = 1
	}
	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	return arr
}

func main() {
	fmt.Println(generate(5))
}
