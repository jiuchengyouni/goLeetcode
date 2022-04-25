package main

//深度优先

type TreeNode23 struct {
	Val   int
	Left  *TreeNode23
	Right *TreeNode23
}

func mergeTrees(t1, t2 *TreeNode23) *TreeNode23 {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
func main() {

}
