package main

//206
//迭代
type ListNode26 struct {
	Val  int
	Next *ListNode26
}

func reverseList(head *ListNode25) *ListNode25 {
	var prev *ListNode25
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {

}
