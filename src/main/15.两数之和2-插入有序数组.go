package main

import "fmt"

//167
func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		if numbers[low]+numbers[high] == target {
			return []int{low, high}
		} else if numbers[low]+numbers[high] > target {
			high--
		} else {
			low++
		}
	}
	return []int{-1, 1}
}

func main() {
	numbers := []int{2, 3, 5, 6, 8}
	fmt.Print(twoSum(numbers, 5))
}
