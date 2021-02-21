package main

type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor1() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	x := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	x := this.s2[len(this.s2)-1]
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

func main() {
	obj := Constructor1()
	obj.Push(1)
	obj.Push(2)
	obj.Pop()
	obj.Peek()
	obj.Empty()

}
