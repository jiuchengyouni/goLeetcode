package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		} else {
			left, right := i+1, len(nums)-1
			for left < right {
				if nums[left]+nums[right] == -nums[i] {
					ans1 := []int{nums[left], nums[right], nums[i]}
					ans = append(ans, ans1)
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < (-nums[i]) {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
