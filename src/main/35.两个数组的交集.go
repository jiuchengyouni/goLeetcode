package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) (ans []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	temp1, temp2, temp3 := 0, 0, -1
	if len(nums2) >= len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	for temp1 < len(nums1) && temp2 < len(nums2) {
		if nums1[temp1] > nums2[temp2] && temp2 != len(nums2)-1 {
			temp2++
			if nums1[temp1] != nums2[temp2] && temp1 != 0 {
				temp1--
			}
		}
		if temp1 > temp3 {
			if nums1[temp1] == nums2[temp2] {
				ans = append(ans, nums1[temp1])
				temp3 = temp1
				temp2++
			}
		}
		temp1++
	}
	return ans
}

func main() {
	nums1 := []int{61, 24, 20, 58, 95, 53, 17, 32, 45, 85, 70, 20, 83, 62, 35, 89, 5, 95, 12, 86, 58, 77, 30, 64, 46, 13, 5, 92, 67, 40, 20, 38, 31, 18, 89, 85, 7, 30, 67, 34, 62, 35, 47, 98, 3, 41, 53, 26, 66, 40, 54, 44, 57, 46, 70, 60, 4, 63, 82, 42, 65, 59, 17, 98, 29, 72, 1, 96, 82, 66, 98, 6, 92, 31, 43, 81, 88, 60, 10, 55, 66, 82, 0, 79, 11, 81}
	nums2 := []int{5, 25, 4, 39, 57, 49, 93, 79, 7, 8, 49, 89, 2, 7, 73, 88, 45, 15, 34, 92, 84, 38, 85, 34, 16, 6, 99, 0, 2, 36, 68, 52, 73, 50, 77, 44, 61, 48}
	fmt.Println(intersect(nums1, nums2))
	nums3 := []int{4, 3, 9, 3, 1, 9, 7, 6, 9, 7}
	nums4 := []int{5, 0, 8}
	fmt.Println(intersect(nums3, nums4))
}
