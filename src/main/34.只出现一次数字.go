package main

import "fmt"

//异或运算

func singleNumber(nums []int) int {
	rdeuce := 0
	for _, v := range nums {
		rdeuce = rdeuce ^ v
	}
	return rdeuce
}

func main() {
	nums := []int{1, 1, 3, 4, 5, 4, 5}
	fmt.Println(singleNumber(nums))
}
