package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var tail *ListNode
	var tail2 *ListNode
	var headtail *ListNode
	var headtail2 *ListNode
	tail = head
	headtail = tail
	tail2 = head.Next
	headtail2 = tail2
	for tail2 != nil && tail2.Next != nil {
		tail.Next = tail.Next.Next
		tail2.Next = tail2.Next.Next
		tail = tail.Next
		tail2 = tail2.Next
	}
	tail.Next = headtail2
	return headtail
}

func main() {

}
