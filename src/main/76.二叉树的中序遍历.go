package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//原因就是append操作会返回这个扩展了的slice的引用，必须让原引用重新赋值为新slice的引用，
//说白了就是，传递过来的这个指针原来指了内存中的A区域，A区域是原数组的真正所在。经过一次 append之后，
//要把这个指针改为指向B，B对应append后新的slice的引用
//但是方法调用的上下文里的slice指针还是指向了老的A内存区域。
//解决办法，方法也很简单，就是传递指针的指针。
func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	inorder(root, &arr)
	return arr
}
func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func main() {
}
