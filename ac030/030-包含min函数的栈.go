package ac030

import (
	"math"
)

/*
题目：
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

直接定义[]int比list.List更快；维护一个辅助栈存放最小值
*/


type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
