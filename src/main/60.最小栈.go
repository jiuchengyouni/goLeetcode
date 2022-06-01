package main

import (
	"fmt"
	"math"
	"sort"
)

type MinStack struct {
	nums     []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		nums:     []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(top, val))
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
