package plusplus

type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.shift()
	ret := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.shift()
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack)+len(this.outStack) == 0
}

func (this *MyQueue) shift() {
	if len(this.outStack) != 0 {
		return
	}
	for len(this.inStack) != 0 {
		tp := this.inStack[len(this.inStack)-1]
		this.inStack = this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack, tp)
	}
}
