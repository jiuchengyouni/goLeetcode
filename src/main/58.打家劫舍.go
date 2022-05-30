package main

import "fmt"

func max1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func rob(nums []int) int {
	var dp = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max1(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max1(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
