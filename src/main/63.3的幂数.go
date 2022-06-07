package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n < 0 || n == 1 {
		return false
	}
	if n%3 == 0 {
		return isPowerOfThree(n / 3)
	} else {
		return true
	}
}

func main() {
	fmt.Println(isPowerOfThree(5))
}
