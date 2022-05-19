package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	length := len(s)
	length1 := 0
	for l := 0; l < length; l++ {
		if unicode.IsLetter(rune(s[l])) || unicode.IsNumber(rune(s[l])) {
			length1++
		}
	}
	fast, low := 0, length-1
	for i := 0; i < length1; i++ {
		for !unicode.IsLetter(rune(s[fast])) && !unicode.IsNumber(rune(s[fast])) {
			fast++
		}
		for !unicode.IsLetter(rune(s[low])) && !unicode.IsNumber(rune(s[low])) {
			low--
		}
		if strings.ToLower(string(s[fast])) != strings.ToLower(string(s[low])) {
			return false
		}
		fast++
		low--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s2 := "race a car"
	fmt.Println(isPalindrome(s2))
	s3 := "0P"
	fmt.Println(isPalindrome(s3))
}
