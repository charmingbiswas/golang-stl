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

func (this *stack[T]) Pop() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}

func (this *stack[T]) Top() (T, bool) {
	if len(this.data) == 0 {
		return *new(T), false
	}

	return this.data[len(this.data)-1], true
}

func (this *stack[T]) IsEmpty() bool {
	return len(this.data) == 0
}
