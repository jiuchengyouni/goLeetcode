package main

import "fmt"

//283
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
func main() {
	nums := []int{0, 1, 2, 0}
	moveZeroes(nums)
	fmt.Print(nums)
}
