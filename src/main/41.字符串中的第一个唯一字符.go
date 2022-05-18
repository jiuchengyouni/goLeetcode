package main

import "fmt"

func firstUniqChar(s string) int {
	length := len(s)
	a := make([]int, 26)
	for i := 0; i < length; i++ {
		a[int(s[i]-97)]++
	}
	for i := 0; i < length; i++ {
		if a[int(s[i]-97)] == 1 {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Print(firstUniqChar("aacd"))
}
