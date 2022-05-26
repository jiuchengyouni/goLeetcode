package main

type TreeNode52 struct {
	Val   int
	Left  *TreeNode52
	Right *TreeNode52
}

func isSymmetric(root *TreeNode52) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode52, right *TreeNode52) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || right.Val != left.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func main() {

}
