package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	arr := []*ListNode{}
	for headA != nil {
		arr = append(arr, headA)
		headA = headA.Next
	}
	for headB != nil {
		if search(arr, headB) {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func search(arr []*ListNode, a *ListNode) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func main() {

}
