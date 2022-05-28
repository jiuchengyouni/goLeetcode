package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums3 []int
	if n == 0 && m == 0 {
		nums1 = nil
		fmt.Println(nums3)
	}
	if n == 0 && m != 0 {
		nums3 = nums1[0:m]
		fmt.Println(nums3)
	}
	if m == 0 && n != 0 {
		nums3 = nums2[0:n]
		fmt.Println(nums3)
	}
	if m != 0 && n != 0 {
		nums3 = nums1[0:m]
		nums3 = append(nums3, nums2[0:n]...)
		sort.Ints(nums3)
		fmt.Println(nums1)
	}
	copy(nums1, nums3)
}

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}
