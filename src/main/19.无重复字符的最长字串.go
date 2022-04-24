package main

import "fmt"

//3
//滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	s := "abcabcd"
	fmt.Println(lengthOfLongestSubstring(s))
}
