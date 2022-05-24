package main

type ListNode49 struct {
	Val  int
	Next *ListNode49
}

func hasCycle(head *ListNode48) bool {
	var map1 map[*ListNode48]int
	map1 = make(map[*ListNode48]int, 100)
	for head != nil {
		map1[head] += 1
		if map1[head] == 2 {
			return true
		} else {
			head = head.Next
		}
	}
	return true
}

func main() {

}
