package main

import "fmt"

func maxProfit(prices []int) int {
	var (
		total int
		index int
	)
	length := len(prices)
	for index < length {
		for index < length-1 && prices[index+1] <= prices[index] {
			index++
		}
		min := prices[index]
		for index < length-1 && prices[index+1] >= prices[index] {
			index++
		}
		total += prices[index] - min
		index++
	}
	return total
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
