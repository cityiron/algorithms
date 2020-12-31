package leetcode

type MyQueue struct {
	in  []int // 输入栈
	out []int // 输出栈
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	r := this.Peek()
	this.out = this.out[:len(this.out)-1]
	return r
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) > 0 {
		return this.out[len(this.out)-1]
	}

	if len(this.out) == 0 && len(this.in) > 0 {
		this.move()
	}

	return this.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in)+len(this.out) == 0
}

// 相当于从 in 出栈，到 out 入栈
func (this *MyQueue) move() {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = []int{}
}
