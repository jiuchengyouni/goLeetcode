package main

import "fmt"

func isAnagram(s string, t string) bool {
	length1 := len(s)
	length2 := len(t)
	if length1 != length2 {
		return false
	}
	a := make([]int, 26)
	for i := 0; i < length1; i++ {
		a[int(s[i]-97)]++
	}
	for i := 0; i < length2; i++ {
		a[int(t[i]-97)]--
	}
	for i := 0; i < 26; i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
