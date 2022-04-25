package main

//116
type Node24 struct {
	Val   int
	Left  *Node24
	Right *Node24
	Next  *Node24
}

func connect(root *Node24) *Node24 {
	if root == nil {
		return root
	}
	for leftmost := root; leftmost != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func main() {

}
