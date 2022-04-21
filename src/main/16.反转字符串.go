package main

import "fmt"

func reverseString(s []byte) {
	first, last := 0, len(s)-1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Print(string(s))
}
