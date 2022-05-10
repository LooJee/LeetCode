package lcof_min_stack

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, this.min(this.minStack[len(this.minStack)-1], x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func (this *MinStack) min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
