package main

import "fmt"

//1672
func maximumWealth(accounts [4][6]int) int {
	max := 0
	for _, account := range accounts {
		m := 0
		for _, acc := range account {
			m += acc
		}
		if m > max {
			max = m
		}
	}
	return max

}
func main() {
	var arr2 [4][6]int
	arr2[1][1] = 1
	arr2[2][2] = 1
	arr2[3][3] = 1
	fmt.Print(maximumWealth(arr2))
}
