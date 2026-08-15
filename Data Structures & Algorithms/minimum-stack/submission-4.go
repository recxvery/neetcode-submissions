type Node struct {
	Val int
	Min int
}

type MinStack struct {
	Stack []Node
}

func Constructor() MinStack {
	return MinStack{
		Stack: make([]Node, 0, 30000),
	}
}

func (this *MinStack) Push(val int) {
	min := val

	if len(this.Stack) > 0 {
		lastmin := this.Stack[len(this.Stack) - 1].Min
		if min > lastmin {
			min = lastmin
		}
	}

	this.Stack = append(this.Stack, Node{Val: val, Min: min})
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.Stack[len(this.Stack) - 1].Min
}

