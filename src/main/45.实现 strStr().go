package main

import "fmt"

func strStr(haystack string, needle string) int {
	lenght1 := len(haystack)
	lenght2 := len(needle)
	for i := 0; i < lenght1-lenght2+1; i++ {
		if haystack[i] == needle[0] {
			str := haystack[i : i+lenght2]
			if str == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}
