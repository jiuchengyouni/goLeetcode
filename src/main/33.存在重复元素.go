package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
