package offer30

import "math"

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	this.stack = this.stack[:l-1]
	this.min = 1000000
	for _, v := range this.stack {
		if this.min > v {
			this.min = v
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
