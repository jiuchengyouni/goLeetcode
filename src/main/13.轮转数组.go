package main

import "fmt"

//189
//数组翻转
func reserve(nums []int) {
	for i, j := 0, len(nums); i < len(nums)/2; i++ {
		nums[i], nums[j-1-i] = nums[j-1-i], nums[i]
	}
}
func rotate(nums []int, k int) {
	k %= len(nums)
	reserve(nums)
	reserve(nums[:k])
	reserve(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 3)
	fmt.Println(nums)
}
