package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 1 {
		return strs[0]
	}
	var length1 int
	var str string
	length1 = len(strs[length-1])
	for i := 0; i < length-1; i++ {
		len1 := len(strs[i])
		if len1 == 0 || length1 == 0 {
			return str
		}
		length1 = min(len1, length1)
	}
	if length1 == 0 {
		return str
	} else {
		for i := 0; i < length1; i++ {
			for j := 0; j < length-1; j++ {
				if strs[j][i] != strs[j+1][i] {
					return str
				}
			}
			if i == 0 {
				str = string(strs[0][0])
			} else {
				str = strs[0][0 : i+1]
			}
		}
	}
	return str
}

func main() {
	str := []string{"caa", "", "a", "acb"}
	fmt.Println(longestCommonPrefix(str))
	str2 := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(str2))
	str3 := []string{"abbb", "a", "accc", "aa"}
	fmt.Println(longestCommonPrefix(str3))
}
