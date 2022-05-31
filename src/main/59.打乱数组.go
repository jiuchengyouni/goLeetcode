package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums     []int
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums,
		append([]int(nil), nums...),
	}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for i := range this.nums {
		j := i + rand.Intn(n-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums
}

func main() {
	nums := []int{1, 2, 3}
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
	var s Solution
	nums = []int{1, 2, 3}
	s = Constructor(nums)
	fmt.Println(s.Shuffle())
}
