package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1]++
	for i := n - 1; i >= 0; i-- {
		if digits[i] > 9 && i != 0 {
			digits[i] -= 10
			digits[i-1]++
		} else if digits[i] > 9 && i == 0 {
			digits[i] -= 10
			ans := []int{}
			ans = append(ans, 1)
			for i := 0; i < n; i++ {
				ans = append(ans, digits[i])
			}
			digits = ans
		}
	}
	return digits
}

func main() {
	digits := []int{1, 9, 9}
	fmt.Println(plusOne(digits))
}
