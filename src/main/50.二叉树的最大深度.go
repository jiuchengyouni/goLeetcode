package main

type TreeNode50 struct {
	Val   int
	Left  *TreeNode50
	Right *TreeNode50
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxDepth(root *TreeNode50) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func main() {

}
