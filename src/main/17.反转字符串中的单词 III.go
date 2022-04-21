package main

import (
	"fmt"
	"strings"
)

//557
//利用String库拆分并反转
func reverseWords(s string) string {
	s1 := strings.Split(s, " ")
	for r := range s1 {
		s1[r] = string(reverse([]byte(s1[r])))
	}
	return strings.Join(s1, " ")
}
func reverse(s []byte) []byte {
	f, e := 0, len(s)-1
	for f < e {
		s[f], s[e] = s[e], s[f]
		f, e = f+1, e-1
	}
	return s
}
func main() {
	s := "Let's take LeetCode contest"
	fmt.Print(reverseWords(s))
}
