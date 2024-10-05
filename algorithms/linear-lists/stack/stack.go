package main

type Stack struct {
	cap   uint64
	depth uint64
	list  []int
}

func NewStack(cap uint64) *Stack {
	return &Stack{
		cap:  cap,
		list: make([]int, cap),
	}
}

func (s *Stack) Push(value int) {
	if s.depth >= s.cap {
		panic("out of capacity.")
	}
	s.list[s.depth] = value
	s.depth++
}

func (s *Stack) Pop() int {
	if s.depth > 0 {
		s.depth--
		return s.list[s.depth]
	}
	return -1
}

func main() {
	s := NewStack(3)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	println(s.Pop())
}
