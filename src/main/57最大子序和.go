package main

import (
	"fmt"
)

func max1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxSubArray(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	max := arr[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = nums[i] + max1(arr[i-1], 0)
		max = max1(max, arr[i])
	}
	return max
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
