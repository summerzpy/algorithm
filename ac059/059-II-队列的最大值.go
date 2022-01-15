package ac059


type MaxQueue struct {
	q []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int)  {
	this.q = append(this.q, value)
	for len(this.max) > 0 && value > this.max[len(this.max) - 1] {
		this.max = this.max[: len(this.max) - 1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.max[0] {
		this.max = this.max[1: ]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}
