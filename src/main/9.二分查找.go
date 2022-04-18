package main

import "fmt"

//704
func search(nums []int, target int) int {
	first := 0
	last := len(nums) - 1
	for first <= last {
		mid := (last-first)/2 + first
		if nums[mid] == target {
			return mid + 1
		} else if nums[mid] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}

	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}
