package main

import "fmt"

//动态规划
func climbStairs(n int) int {
	q, temp1, temp2 := 1, 0, 0
	for i := 0; i < n; i++ {
		temp2 = temp1
		temp1 = q
		q = temp1 + temp2
	}
	return q
}
func main() {
	fmt.Println(climbStairs(3))
}
