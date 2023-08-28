package main

type MyStack struct {
	elms []int
}

func Constructor225() MyStack {
	return MyStack{
		elms: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.elms = append(this.elms, x)
}

func (this *MyStack) Pop() int {
	v := this.elms[len(this.elms)-1]
	this.elms = this.elms[:len(this.elms)-1]
	return v
}

func (this *MyStack) Top() int {
	return this.elms[len(this.elms)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.elms) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
