package main

import "fmt"

//1342
func numberOfSteps(num int) int {
	var sum = 0
	for {
		if num == 0 {
			break
		} else {
			if (num % 2) == 0 {
				num /= 2
				sum += 1
			} else {
				num -= 1
				sum += 1
			}
		}
	}
	return sum
}
func main() {
	fmt.Println(numberOfSteps(8))
}
