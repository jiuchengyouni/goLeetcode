package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		//右移运算符>>,运算结果正好能对应一个整数的二分之一值，这就正好能代替数学上的除2运算，但是比除2运算要快。
		//mid=(left+right)>>1相当于mid=(left+right)/2
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
func main() {

}
