package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	n := len(s)
	sign := 1
	var (
		i   int
		abs int
	)
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = abs*10 + int(s[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return abs * sign
}

func main() {
	s := "4193 with words"
	fmt.Println(myAtoi(s))
}
