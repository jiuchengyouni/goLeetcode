package main

import (
	"fmt"
	"math"
)

type TreeNode51 struct {
	Val   int
	Left  *TreeNode51
	Right *TreeNode51
}

func isValidBST(root *TreeNode51) bool {
	stack := []*TreeNode51{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	fmt.Println(stack[1])
	inorder := math.MaxInt64
	fmt.Println(inorder)
}
