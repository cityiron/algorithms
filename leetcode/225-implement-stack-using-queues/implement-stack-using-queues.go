package leetcode

// https://leetcode-cn.com/problems/implement-stack-using-queues/

// 两个队列实现

type MyStack struct {
	q1 []int // 这里对应stack数据
	q2 []int // 辅助的
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: []int{},
		q2: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x) // 加到 q2
	// 把 q1 的内容移动到 q2
	for i := 0; i < len(this.q1); i++ {
		this.q2 = append(this.q2, this.q1[i])
	}
	this.q1, this.q2 = this.q2, []int{}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	r := this.Top()
	this.q1 = this.q1[1:]
	return r
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

// 一个数组的实现

type MyStack1 struct {
	element []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack1 {
	return MyStack1{
		element: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack1) Push(x int) {
	this.element = append(this.element, x)
	if len(this.element) > 1 {
		// 拿出第一个，后面的往后移动
		front := this.element[len(this.element)-1]
		for i := len(this.element) - 1; i > 0; i-- {
			this.element[i] = this.element[i-1]
		}
		this.element[0] = front
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack1) Pop() int {
	r := this.Top()
	this.element = this.element[1:]
	return r
}

/** Get the top element. */
func (this *MyStack1) Top() int {
	return this.element[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack1) Empty() bool {
	return len(this.element) == 0
}

// 直接数组实现，只是用来参考对比用

type MyStack2 struct {
	element []int
}

/** Initialize your data structure here. */
func Constructor2() MyStack2 {
	return MyStack2{
		element: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack2) Push(x int) {
	this.element = append(this.element, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack2) Pop() int {
	r := this.Top()
	this.element = this.element[:len(this.element)-1]
	return r
}

/** Get the top element. */
func (this *MyStack2) Top() int {
	return this.element[len(this.element)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack2) Empty() bool {
	if this.element != nil && len(this.element) > 0 {
		return false
	}

	return true
}
