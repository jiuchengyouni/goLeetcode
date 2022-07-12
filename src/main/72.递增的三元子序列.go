package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	//记录最小值和第二小的值
	m1, m2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		//找到子序列第一个元素，不断更新
		if m1 >= v {
			m1 = v
		} else if m2 >= v {
			//找到子序列第二个元素，不断更新
			m2 = v
		} else {
			//找到第三个，满足要求
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2}
	fmt.Println(increasingTriplet(nums))

}
