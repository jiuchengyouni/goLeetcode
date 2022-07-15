package main

type ListNode73 struct {
	Val  int
	Next *ListNode73
}

func addTwoNumbers(l1 *ListNode73, l2 *ListNode73) *ListNode73 {
	var tail *ListNode73
	var head *ListNode73
	temp := 0
	for l1 != nil || l2 != nil {
		n1 := 0
		n2 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + temp
		sum, temp = sum%10, sum/10
		if head == nil {
			head = &ListNode73{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode73{Val: sum}
			tail = tail.Next
		}
		if temp != 0 {
			tail.Next = &ListNode73{Val: temp}
		}
	}
	return head
}

func main() {

}
