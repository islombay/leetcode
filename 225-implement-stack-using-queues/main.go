package main

import "container/list"

type MyStack struct {
	l *list.List
}

func Constructor() MyStack {
	return MyStack{l: list.New()}
}

func (this *MyStack) Push(x int) {
	this.l.PushBack(x)
}

func (this *MyStack) Pop() int {
	return this.l.Remove(this.l.Back()).(int)
}

func (this *MyStack) Top() int {
	return this.l.Back().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.l.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
