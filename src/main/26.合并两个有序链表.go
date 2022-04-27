package main

//21
type ListNode25 struct {
	Val  int
	Next *ListNode25
}

func mergeTwoLists(list1 *ListNode25, list2 *ListNode25) *ListNode25 {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
func main() {

}
