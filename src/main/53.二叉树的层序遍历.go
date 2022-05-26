package main

type TreeNode53 struct {
	Val   int
	Left  *TreeNode53
	Right *TreeNode53
}

func levelOrder(root *TreeNode53) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode53{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode53{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {

}
