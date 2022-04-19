package main

import "fmt"

//977
//双指针
func sortedSquares(nums []int) []int {
	len := len(nums)
	first := 0
	last := len - 1
	ans := make([]int, len)
	for i := len - 1; i >= 0; i-- {
		v := nums[last] * nums[last]
		w := nums[first] * nums[first]
		if v > w {
			ans[i] = v
			last--
		} else {
			ans[i] = w
			first++
		}
	}
	return ans
}
func main() {
	nums := []int{-1, 2, 4}
	fmt.Println(sortedSquares(nums))
}
