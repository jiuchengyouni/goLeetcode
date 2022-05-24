package main

type ListNode48 struct {
	Val  int
	Next *ListNode48
}

func deleteNode(node *ListNode48) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {

}
