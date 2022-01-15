package ac009

import "container/list"

/*
题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead 操作返回 -1 )
*/
type CQueue2 struct {
	// 借助list包实现，主要借助PushBack, Remove, Back， Len实现
	//stack1用于入队，stack2用于出队
	stack1, stack2 *list.List
}

func Constructor2() CQueue2 {
	return CQueue2{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue2) AppendTail2(val int)  {
	// 添加队尾-入队
	this.stack1.PushBack(val)
}

func (this *CQueue2) DeleteHead2() int {
	// 删除队头-出队
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}




