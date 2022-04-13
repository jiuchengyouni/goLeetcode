package main

import (
	"fmt"
	"strconv"
)

//字符拼接
func fizzBuzz(n int) (arr []string) {
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if (i+1)%3 == 0 {
			arr = append(arr, "Fizz")
		} else if (i+1)%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i+1))
		}
	}
	return
}

func main() {
	fmt.Print(fizzBuzz(5))
}
