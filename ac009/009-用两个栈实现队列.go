package ac009

type CQueue struct {
	in stack
	out stack
}

type stack []int

func(s *stack) Push(value int) {
	*s = append(*s, value)
}

func(s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n - 1]

	*s = (*s)[:n - 1]

	return res
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.in.Push(value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.out) != 0 {
		return this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}

	return -1

}

