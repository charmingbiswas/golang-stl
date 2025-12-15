package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, 10),
	}
}

func (this *Stack[T]) Push(val T) {
	this.data = append(this.data, val)
}

func (this *Stack[T]) Pop() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *Stack[T]) Top() T {
	if len(this.data) == 0 {
		return *new(T)
	}

	return this.data[len(this.data)-1]
}

func (this *Stack[T]) IsEmpty() bool {
	return len(this.data) == 0
}
