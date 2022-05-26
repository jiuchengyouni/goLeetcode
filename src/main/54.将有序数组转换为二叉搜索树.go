package main

type TreeNode54 struct {
	Val   int
	Left  *TreeNode54
	Right *TreeNode54
}

func sortedArrayToBST(nums []int) *TreeNode54 {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBST2(nums, 0, len(nums)-1)
}

func sortedArrayToBST2(nums []int, start int, end int) *TreeNode54 {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	var root TreeNode54
	root.Val = nums[mid]
	root.Left = sortedArrayToBST2(nums, start, mid-1)
	root.Right = sortedArrayToBST2(nums, mid+1, end)
	return &root
}

func main() {

}
