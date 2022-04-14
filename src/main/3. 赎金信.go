package main

import "fmt"

//383
//字符统计
func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	//ascii码
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Print(canConstruct("aaa", "aab"))
}
