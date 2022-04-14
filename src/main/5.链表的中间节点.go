package main

//876
//快慢指针

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func middleNode(head *ListNode2) *ListNode2 {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func main() {

}
