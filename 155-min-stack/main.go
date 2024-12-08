package main

type node struct {
	Val     int
	Next    *node
	NextMin *node
}

type MinStack struct {
	head *node
	min  *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	newNode := &node{Val: val}
	if this.head != nil {
		newNode.Next = this.head
	}
	this.head = newNode
	if this.min == nil || this.min.Val > val {
		newNode.NextMin = this.min
		this.min = newNode
	}
}

func (this *MinStack) Pop() {
	if this.min == this.head {
		this.min = this.min.NextMin
	}
	this.head = this.head.Next
}

func (this *MinStack) Top() int {
	return this.head.Val
}

func (this *MinStack) GetMin() int {
	return this.min.Val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
