package queue_by_two_stack

type CQueue struct {
	input  []int
	output []int
}

func Constructor() CQueue {
	return CQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	if this == nil {
		return
	}

	this.input = append(this.input, value)
}

func (this *CQueue) DeleteHead() int {
	if this == nil {
		return -1
	}

	if len(this.output) == 0 {
		if len(this.input) == 0 {
			return -1
		}

		this.move()
	}

	v := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]

	return v
}

func (this *CQueue) move() {
	if this == nil {
		return
	}

	for len(this.input) != 0 {
		this.output = append(this.output, this.input[len(this.input)-1])
		this.input = this.input[:len(this.input)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
