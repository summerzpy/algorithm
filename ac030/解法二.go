package ac030

import "container/list"

type MinStack1 struct {
	stack *list.List
	minStack *list.List
}


/** initialize your data structure here. */
func Constructor1() MinStack1 {
	return MinStack1{
		stack: list.New(),
		minStack: list.New(),
	}
}


func (this *MinStack1) Push1(x int)  {
	this.stack.PushBack(x)
	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		this.minStack.PushBack(min(this.minStack.Back().Value.(int),x))
	}
}


func (this *MinStack1) Pop1()  {
	this.stack.Remove(this.stack.Back())
	this.minStack.Remove(this.minStack.Back())
}


func (this *MinStack1) Top1() int {
	return this.stack.Back().Value.(int)
}


func (this *MinStack1) Min1() int {
	return this.minStack.Back().Value.(int)
}

func min1(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */