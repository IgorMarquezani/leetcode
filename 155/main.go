package main

import "fmt"

type Node struct {
	v     int
	clone *int
}

type MinStack struct {
	buf    []Node
	sorted []int
}

func Constructor() MinStack {
	return MinStack{
		buf:    make([]Node, 0, 64),
		sorted: make([]int, 0, 64),
	}
}

func (this *MinStack) Push(val int) {
	this.buf = append(this.buf, Node{v: val})

	if len(this.sorted) == 0 {
		this.sorted = append(this.sorted, val)
		return
	}
	for i := range this.sorted {
		if val < this.sorted[i] {
			tmp := make([]int, len(this.sorted[:i]))

			copy(tmp, this.sorted[:i])
			tmp = append(tmp, val)
			tmp = append(tmp, this.sorted[i:]...)

			this.sorted = tmp

			break
		}
	}
	if len(this.buf) > len(this.sorted) {
		this.sorted = append(this.sorted, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.buf) > 0 {
		this.buf = this.buf[:len(this.buf)-1]
		this.sorted = this.sorted[:len(this.sorted)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.buf) > 0 {
		return this.buf[len(this.buf)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.sorted) > 0 {
		return this.sorted[0]
	}
	return 0
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.sorted)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())

	fmt.Println(s.buf)
	fmt.Println(s.sorted)
}
