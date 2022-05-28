package main

import "math"

func max1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxProfit(prices []int) int {
	var (
		total int
		index int
	)
	min := math.MaxInt64
	max := 0
	length := len(prices)
	for index < length {
		for index < length-1 && prices[index+1] <= prices[index] {
			index++
		}
		if min > prices[index] {
			min = prices[index]
			max = 0
		}
		for index < length-1 && prices[index+1] >= prices[index] {
			index++
		}
		if max < prices[index] {
			max = prices[index]
		}
		if max-min > 0 {
			total = max1(total, max-min)
		}
		index++
	}
	return total
}

func main() {

}
