package main

import "fmt"

func missingNumber(nums []int) int {
	length := len(nums)
	a := make(map[int]int, length+1)
	for i := 0; i < length+1; i++ {
		a[i] = 0
	}
	for i := 0; i < length; i++ {
		a[nums[i]] = 1
	}
	for i := 0; i < length; i++ {
		if a[i] == 0 {
			return i
		}
	}
	return length
}

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
}
