package main

import "fmt"

func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	num -= 1
	num -= 1
	fmt.Println(num)
}
