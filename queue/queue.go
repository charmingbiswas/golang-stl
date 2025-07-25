package queue

type queue[T any] struct {
	data []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		data: make([]T, 0, 10),
	}
}

func (this *queue[T]) Push(val T) {
	this.data = append(this.data, val)
}

func (this *queue[T]) Pop() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

func (this *queue[T]) Front() (T, bool) {
	if len(this.data) == 0 {
		return *new(T), false
	}

	return this.data[0], true
}

func (this *queue[T]) IsEmpty() bool {
	return len(this.data) == 0
}
