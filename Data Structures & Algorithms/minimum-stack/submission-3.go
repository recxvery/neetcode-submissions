type MinStack struct {
	Stack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Stack) == 0 {
		return 0
	}

	if len(this.Stack) == 1 {
		return this.Stack[0]
	}
	
	min := this.Stack[0]

	for i := 1; i < len(this.Stack); i++ {
		if this.Stack[i] < min {
			min = this.Stack[i]
		}
	}

	return min
}

