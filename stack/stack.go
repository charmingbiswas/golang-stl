package stack

type stack[T any] struct {
	data []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		data: make([]T, 0, 10),
	}
}

func (this *stack[T]) Push(val T) {
	this.data = append(this.data, val)
}

func (this *stack[T]) Pop() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *stack[T]) Top() T {
	if len(this.data) == 0 {
		return *new(T)
	}

	return this.data[len(this.data)-1]
}

func (this *stack[T]) IsEmpty() bool {
	return len(this.data) == 0
}
