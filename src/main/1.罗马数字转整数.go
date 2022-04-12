package main

import (
	"fmt"
)

//013
//哈希表+遍历，从左至右或者从右至左遍历都可以

var value = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	n := len(s)
	for i := 0; i < n; i++ {
		valuepart := value[s[i]]
		if i < n-1 && valuepart < value[s[i+1]] {
			sum -= valuepart
		} else {
			sum += valuepart
		}
	}
	return sum
}
func main() {
	fmt.Print(romanToInt("LVIII"))
	fmt.Print(romanToInt("I"))
}
