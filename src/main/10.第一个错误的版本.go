package main

import "sort"

func firstBadVersion(n int) int {
	i := 1
	//队首
	for i < n {
		mid := i + (n-i)/2
		if isBadVersion(mid) {
			n = mid
		} else {
			i = mid + 1
		}
	}
	return n
}

//利用sort
func firstBadVersion2(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
func main() {

}
