package main

import "fmt"

func removeDuplicates(nums []int) int {
	fast, low := 0, 0
	length := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			fast++
		} else {
			length++
			low++
			nums[low] = nums[fast+1]
			fast++
		}
	}
	return length
}
func main() {
	ans := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(ans))
	fmt.Println(ans)
}
