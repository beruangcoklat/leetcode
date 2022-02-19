type node struct {
	val int
	min int
}

type MinStack struct {
	stack  []node
	length int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	n := node{
		val: val,
		min: func() int {
			lastIdx := this.length - 1
			if this.length == 0 || val < this.stack[lastIdx].min {
				return val
			}
			return this.stack[lastIdx].min
		}(),
	}
	this.stack = append(this.stack, n)
	this.length++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[this.length-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
